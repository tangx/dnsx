package aliyun

import (
	"errors"
	"log"
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

// Delete record
func (ali AliyunDNS) Delete(subDomain string) {
	cli := ali.InitClient()
	r, err := ali.SubDomainInfo(subDomain)
	if err != nil {
		log.Fatal(err)
	}

	_, errResp, err := cli.DeleteDomainRecord(r.RecordId)
	if err != nil {
		log.Fatal(errResp.Message)
	}
	logrus.Infof("Deleted: %s.%s (%s %s)", subDomain, ali.Domain, r.Type, r.Value)

}

// SubDomainInfo return records info
func (ali AliyunDNS) SubDomainInfo(subDommain string) (alidns.RecordInfo, error) {
	cli := ali.InitClient()
	data := map[string]string{
		"PageSize": "500",
	}
	resp, errResp, err := cli.DescribeDomainRecords(ali.Domain, data)
	if err != nil {
		logrus.Fatal(errResp.Message)
	}

	for _, record := range resp.DomainRecords.Record {
		if record.RR == subDommain {
			return record, nil
		}
	}
	return alidns.RecordInfo{}, errors.New("subdomain not found")
}
