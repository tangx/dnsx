package aliyun

import (
	"github.com/sirupsen/logrus"
	"github.com/tangx/alidns-sdk"
)

// AliyunDNS is cns
type AliyunDNS struct {
	AKID string
	AKEY string
}

// InitClient for aliyun dns
func (ali AliyunDNS) InitClient() *alidns.Client {
	return alidns.New(ali.AKID, ali.AKEY)
}

// Add domain record
func (ali AliyunDNS) Add(domain, rr, rrType, rrValue string) string {
	cli := ali.InitClient()
	resp, errResp, err := cli.AddDomainRecord(domain, rr, rrType, rrValue, nil)
	if err != nil {
		logrus.Fatal(errResp.Message)
	}

	logrus.Infof("Added: %s.%s (%s %s)", rr, domain, rrType, rrValue)
	return resp.RecordId
}
