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

// GetRecords 查询 DNS 解析记录
func (cli Client) GetRecords(domain, record string) (RRs []backend.RecordItem) {
	aliyundns := alidns.New(cli.AKID, cli.AKEY)

	reqBody := map[string]string{"RRKeyWord": record, "PageSize": "500"}

	respInfo, errBody, err := aliyundns.DescribeDomainRecords(domain, reqBody)
	if err != nil {
		logrus.Fatalln(errBody)
	}

	for _, rr := range respInfo.DomainRecords.Record {
		item := backend.RecordItem{
			ID:       rr.RecordId,
			Name:     rr.RR,
			Type:     rr.Type,
			Value:    rr.Value,
			Status:   rr.Status,
			UpdateOn: "", // aliyun 没有记录解析修改时间
		}
		RRs = append(RRs, item)
	}

	return
}

// DeleteRecord 删除域名解析记录
func (cli Client) DeleteRecord(domain string, id string) string {
	aliyundns := alidns.New(cli.AKID, cli.AKEY)
	respBody, errBody, err := aliyundns.DeleteDomainRecord(id)
	if err != nil {
		logrus.Errorf("%s : %s", errBody.RequestID, errBody.Message)
	}

	return respBody.RecordId
}

func (cli Client) SetRecordStatus(domain string, recordID string, status bool) string {
	aliyundns := alidns.New(cli.AKID, cli.AKEY)

	respBody, errBody, err := aliyundns.SetDomainRecordStatus(recordID, status)
	if err != nil {
		logrus.Errorf("%s : %s", errBody.RequestID, errBody.Message)
	}

	return respBody.RecordId
}
