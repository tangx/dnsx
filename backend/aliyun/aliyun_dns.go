package aliyun

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
	"github.com/tangx/alidns-sdk"
	"github.com/tangx/dnsx/backend"
)

// Client 阿里云 DNS
type Client struct {
	AKID string `json:"akid"`
	AKEY string `json:"akey"`
}

// AddRecord 增加解析记录
func (cli Client) AddRecord(domain, record, rrType, Value string) (recordID string) {
	aliyundns := alidns.New(cli.AKID, cli.AKEY)

	respBody, errBody, err := aliyundns.AddDomainRecord(domain, record, rrType, Value, nil)
	if err != nil {
		errBytes, _ := json.Marshal(errBody)
		logrus.Fatalf("%s", errBytes)

	}
	return respBody.RecordId
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
	aliyundns := alidns.New(cli.AKID, cli.AKEY)

	reqBody := map[string]string{"RRKeyWord": record, "PageSize": "500"}

	respInfo, errBody, err := aliyundns.DescribeDomainRecords(domain, reqBody)
	if err != nil {
		logrus.Fatalln(errBody)
	}

	for _, rr := range respInfo.DomainRecords.Record {
		RRs = append(RRs, backend.RecordItem{
			rr.RecordId,
			rr.RR,
			rr.Type,
			rr.Value,
			rr.Status,
		})
	}

	return
}
