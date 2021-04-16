package aliyun

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
	"github.com/tangx/alidns-sdk"
	"github.com/tangx/dnsx/pkg/backend/internal/response"
)

// Client 阿里云 DNS
type Client struct {
	cli *alidns.Client
}

// NewClient 创建一个 aliyun dns 客户端
func NewClient(akid string, akey string) Client {
	// c := Client{AKID: akid, AKEY: akey}
	c := Client{}
	if c.cli == nil {
		// cli :=
		c.cli = alidns.New(akid, akey)
	}
	return c
}

// AddRecord 增加解析记录
func (c Client) AddRecord(domain, record, rrType, Value string) (recordID string) {

	respBody, errBody, err := c.cli.AddDomainRecord(domain, record, rrType, Value, nil)
	if err != nil {
		errBytes, _ := json.Marshal(errBody)
		logrus.Fatalf("%s", errBytes)

	}
	return respBody.RecordId
}

// GetRecords 查询 DNS 解析记录
func (c Client) GetRecords(domain, record string) (RRs []response.RecordItem) {

	reqBody := map[string]string{"RRKeyWord": record, "PageSize": "500"}

	respInfo, errBody, err := c.cli.DescribeDomainRecords(domain, reqBody)
	if err != nil {
		logrus.Fatalln(errBody)
	}

	for _, rr := range respInfo.DomainRecords.Record {
		RRs = append(RRs, response.RecordItem{
			ID:       rr.RecordId,
			Name:     rr.RR,
			Type:     rr.Type,
			Value:    rr.Value,
			Status:   rr.Status,
			UpdateOn: "", // aliyun 没有记录解析修改时间
		})
	}

	return
}

// DeleteRecord 删除域名解析记录
func (c Client) DeleteRecord(domain string, id string) string {

	respBody, errBody, err := c.cli.DeleteDomainRecord(id)
	if err != nil {
		logrus.Errorf("%s : %s", errBody.RequestID, errBody.Message)
	}

	return respBody.RecordId
}

func (c Client) SetRecordStatus(domain string, recordID string, status bool) string {

	respBody, errBody, err := c.cli.SetDomainRecordStatus(recordID, status)
	if err != nil {
		logrus.Errorf("%s : %s", errBody.RequestID, errBody.Message)
	}

	return respBody.RecordId
}
