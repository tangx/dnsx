package qcloud

import (
	"fmt"
	"strings"

	"github.com/tangx/dnsx/config"

	cns "github.com/go-http/qcloud-cns"
	"github.com/sirupsen/logrus"
)

// QcloudCNS is cns
type QcloudCNS struct {
	AKID   string
	AKEY   string
	Domain string
	//cli  *cns.Client
}

func (qcns *QcloudCNS) InitClient() (cli *cns.Client) {
	return cns.New(qcns.AKID, qcns.AKEY)
}

func (qcns QcloudCNS) allRecords() (records []cns.Record) {
	cli := qcns.InitClient()

	records, err := cli.RecordList(qcns.Domain)
	if err != nil {
		logrus.Error(err)
	}
	return
}

func (qcns QcloudCNS) List() (records []cns.Record) {
	records = qcns.allRecords()

	for _, r := range records {
		fmt.Printf("%d %s.%s %s %s\n", r.Id, r.Name, qcns.Domain, r.Type, r.Value)
	}
	return
}

func (qcns QcloudCNS) SubDomainInfo(subDomain string) cns.Record {
	records := qcns.allRecords()

	for _, r := range records {
		if r.Name == subDomain {
			return r
		}
	}
	return cns.Record{}
}

func (qcns QcloudCNS) add(subDomain string, recordType string, recordValue string) (recordID int) {
	cli := qcns.InitClient()

	record := cns.Record{
		Name:  subDomain,
		Type:  recordType,
		Value: recordValue,
	}

	recordID, err := cli.RecordCreate(qcns.Domain, record)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("Added: %s.%s (%s %s)", subDomain, qcns.Domain, recordType, recordValue)

	return
}

func (qcns QcloudCNS) Add(domain, sub, recordType, recordValue string) (recordID int) {
	_, akid, akey := config.LoadDomainConfig(domain)
	qcns.AKID = akid
	qcns.AKEY = akey
	qcns.Domain = domain
	return qcns.add(sub, strings.ToUpper(recordType), recordValue)
}

func (qcns QcloudCNS) Delete(subDomain string) {
	cli := qcns.InitClient()

	r := qcns.SubDomainInfo(subDomain)
	if r.Id == 0 {
		logrus.Fatalf("%s.%s is not Exist", subDomain, qcns.Domain)
	}

	err := cli.RecordDelete(qcns.Domain, r.Id)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Infof("Deleted: %s.%s (%s %s)", subDomain, qcns.Domain, r.Type, r.Value)
}

func (qcns QcloudCNS) Get(subDomain string) {
	record := qcns.SubDomainInfo(subDomain)

	fmt.Println(record.Id, record.Name, record.Type, record.Value)
}

func (qcns QcloudCNS) Update(subDomain string, recordType string, recordValue string) {
	cli := qcns.InitClient()

	record := qcns.SubDomainInfo(subDomain)
	if record.Id == 0 {
		logrus.Fatalf("%s.%s is not Exists", subDomain, qcns.Domain)
	}

	recordOldtype := record.Type
	recordOldValue := record.Value

	record.Type = recordType
	record.Value = recordValue

	err := cli.RecordModify(qcns.Domain, record)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("Updated: %s.%s (%s %s) <- (%s %s)", subDomain, qcns.Domain,
		recordType, recordValue,
		recordOldtype, recordOldValue)

}
