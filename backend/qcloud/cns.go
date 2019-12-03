package qcloud

import (
	cns "github.com/go-http/qcloud-cns"
	"github.com/sirupsen/logrus"
)

// QcloudCNS is cns
type QcloudCNS struct {
	AKID string
	AKEY string
}

// InitClient for QcloudCNS
func (qcns *QcloudCNS) InitClient() (cli *cns.Client) {
	return cns.New(qcns.AKID, qcns.AKEY)
}

// Add Domain Record
func (qcns QcloudCNS) Add(domain, rr, rrType, rrValue string) string {
	cli := qcns.InitClient()

	record := cns.Record{
		Name:  rr,
		Type:  rrType,
		Value: rrValue,
	}

	rrID, err := cli.RecordCreate(domain, record)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("Added: %s.%s (%s %s)", rr, domain, rrType, rrValue)

	return string(rrID)
}
