package aliyun

import (
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/tangx/alidns-sdk"
)

// AliyunDNS is cns
type AliyunDNS struct {
	AKID   string
	AKEY   string
	Domain string
}

// InitClient 初始化dns
func (ali AliyunDNS) InitClient() *alidns.Client {
	return alidns.New(ali.AKID, ali.AKEY)
}

// Add domain record
func (ali AliyunDNS) Add(subDomain string, recordType string, recordValue string) (recordID int) {
	cli := ali.InitClient()
	resp, errResp, err := cli.AddDomainRecord(ali.Domain, subDomain, recordType, recordValue, nil)
	if err != nil {
		logrus.Fatal(errResp.Message)
	}
	rrID, _ := strconv.Atoi(resp.RecordId)

	logrus.Infof("Added: %s.%s (%s %s)", subDomain, ali.Domain, recordType, recordValue)
	return rrID
}
