package qcloud

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	cns "github.com/go-http/qcloud-cns"
	"github.com/sirupsen/logrus"
	"github.com/tangx/dnsx/backend"
)

// Client 腾讯云 DNS
type Client struct {
	AKID string
	AKEY string
}

// AddRecord 添加解析记录
func (cli Client) AddRecord(domain, rr, rrType, rrValue string) (recordID string) {
	qcns := cns.New(cli.AKID, cli.AKEY)
	id, err := qcns.RecordCreate(
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

// // RecordItem response
// type RecordItem struct {
// 	ID     string
// 	Name   string
// 	Type   string
// 	Value  string
// 	Status string
// }

// GetRecords 查询 DNS 解析记录
func (cli Client) GetRecords(domain, record string) (RRs []backend.RecordItem) {
	qcns := cns.New(cli.AKID, cli.AKEY)
	records, err := qcns.RecordList(domain)
	if err != nil {
		logrus.Fatalln(err)
	}

	patt := fmt.Sprintf(".*%s.*", record)
	compl := regexp.MustCompile(patt)

	for _, rr := range records {
		if compl.Match([]byte(rr.Name)) {
			// 偷懒初始化值的警告
			// https://www.maodapeng.com/topic/10030.html
			// composite literal uses unkeyed fields
			RRs = append(RRs, backend.RecordItem{
				strconv.Itoa(rr.Id),
				rr.Name,
				rr.Type,
				rr.Value,
				rr.Status,
			})
		}
	}

	return
}
