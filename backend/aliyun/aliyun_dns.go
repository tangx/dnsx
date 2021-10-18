package aliyun

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
	"github.com/tangx/alidns-sdk"
	"github.com/tangx/dnsx/backend"
)

// Client 阿里云 DNS
type Client struct {
	AKID  string `json:"akid"`
	AKEY  string `json:"akey"`
	agent *alidns.Client
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
		cli.agent = alidns.New(cli.AKID, cli.AKEY)
	}
}

// AddRecord 增加解析记录
func (cli *Client) AddRecord(domain, record, rrType, Value string) (recordID string) {

	respBody, errBody, err := cli.agent.AddDomainRecord(domain, record, rrType, Value, nil)
	if err != nil {
		errBytes, _ := json.Marshal(errBody)
		logrus.Fatalf("%s", errBytes)

	}
	return respBody.RecordId
}

// GetRecords 查询 DNS 解析记录
func (cli *Client) GetRecords(domain, record string) (RRs []backend.RecordItem) {

	reqBody := map[string]string{"RRKeyWord": record, "PageSize": "500"}

	respInfo, errBody, err := cli.agent.DescribeDomainRecords(domain, reqBody)
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
func (cli *Client) DeleteRecord(domain string, id string) string {
	respBody, errBody, err := cli.agent.DeleteDomainRecord(id)
	if err != nil {
		logrus.Errorf("%s : %s", errBody.RequestID, errBody.Message)
	}

	return respBody.RecordId
}

func (cli *Client) SetRecordStatus(domain string, recordID string, status bool) string {

	respBody, errBody, err := cli.agent.SetDomainRecordStatus(recordID, status)
	if err != nil {
		logrus.Errorf("%s : %s", errBody.RequestID, errBody.Message)
	}

	return respBody.RecordId
}
