package aliyun

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/tangx/alidns-sdk"
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
