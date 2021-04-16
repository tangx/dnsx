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
	cli *cns.Client
}

// NewClient 返回一个新的 Qcloud / Dnspod 客户端
func NewClient(akid string, token string) Client {
	c := Client{}
	if c.cli == nil {
		c.cli = cns.New(akid, token)
	}
	return c
}

// AddRecord 添加解析记录
func (c Client) AddRecord(domain, rr, rrType, rrValue string) (recordID string) {
	id, err := c.cli.RecordCreate(
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
func (c Client) GetRecords(domain, record string) (RRs []backend.RecordItem) {

	records, err := c.cli.RecordList(domain)
	if err != nil {
		logrus.Fatalln(err)
	}

	pattern := fmt.Sprintf(".*%s.*", record)
	re := regexp.MustCompile(pattern)

	for _, rr := range records {
		if re.Match([]byte(rr.Name)) {
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
func (c Client) DeleteRecord(domain string, recordID string) string {

	id, _ := strconv.Atoi(recordID)
	err := c.cli.RecordDelete(domain, id)
	if err != nil {
		logrus.Errorf("%s", err)
	}
	return recordID

}

// SetRecordStatus 设置域名解析记录状态
func (c Client) SetRecordStatus(domain string, recordID string, status bool) string {

	rID, _ := strconv.Atoi(recordID)
	err := c.cli.RecordStatus(domain, rID, status)
	if err != nil {
		logrus.Errorf("%s", err)
	}
	return recordID

}
