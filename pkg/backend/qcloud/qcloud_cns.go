package qcloud

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	cns "github.com/go-http/qcloud-cns"
	"github.com/sirupsen/logrus"
	"github.com/tangx/dnsx/pkg/backend"
)

// Client 腾讯云 DNS
type Client struct {
	AKID  string
	AKEY  string
	agent *cns.Client
}

func NewClient(akid, akey string) *Client {
	c := &Client{
		AKID: akid,
		AKEY: akey,
	}

	c.initial()
	return c
}

func (cli *Client) initial() {
	if cli.agent == nil {
		cli.agent = cns.New(cli.AKID, cli.AKEY)
	}
}

// AddRecord 添加解析记录
func (cli *Client) AddRecord(domain, rr, rrType, rrValue string) (recordID string) {
	id, err := cli.agent.RecordCreate(
		domain,
		cns.Record{
			Name:  rr,
			Type:  strings.ToUpper(rrType),
			Value: rrValue,
		})

	if err != nil {
		logrus.Fatalln(err)
	}

	return strconv.Itoa(id)
}

// GetRecords 查询 DNS 解析记录
func (cli *Client) GetRecords(domain, record string) (RRs []backend.RecordItem) {
	records, err := cli.agent.RecordList(domain)
	if err != nil {
		logrus.Fatalln(err)
	}

	// 解决通配符的问题
	record = strings.ReplaceAll(record, "*", `\\*`)
	pattern := fmt.Sprintf(".*%s.*", record)
	re := regexp.MustCompile(pattern)

	for _, rr := range records {
		if re.Match([]byte(rr.Name)) {
			// 偷懒初始化值的警告
			// https://www.maodapeng.com/topic/10030.html
			// composite literal uses unkeyed fields
			var Status string
			if rr.Enabled == 1 {
				Status = "enable"
			} else {
				Status = "disable"
			}

			RRs = append(RRs, backend.RecordItem{
				ID:       strconv.Itoa(rr.Id),
				Name:     rr.Name,
				Type:     rr.Type,
				Value:    rr.Value,
				Status:   Status,
				UpdateOn: rr.UpdatedOn,
			})
		}
	}

	return
}

// DeleteRecord 删除解析记录
func (cli *Client) DeleteRecord(domain string, recordID string) string {

	id, _ := strconv.Atoi(recordID)
	err := cli.agent.RecordDelete(domain, id)
	if err != nil {
		logrus.Errorf("%s", err)
	}
	return recordID
}

func (cli *Client) SetRecordStatus(domain string, recordID string, status bool) string {

	rID, _ := strconv.Atoi(recordID)
	err := cli.agent.RecordStatus(domain, rID, status)
	if err != nil {
		logrus.Errorf("%s", err)
	}
	return recordID
}
