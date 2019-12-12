package qcloud

import (
	"strconv"
	"strings"

	cns "github.com/go-http/qcloud-cns"
	"github.com/sirupsen/logrus"
)

// Client 腾讯云 DNS
type Client struct {
	AKID string
	AKEY string
}

// AddRecord 添加解析记录
func (cli Client) AddRecord(domain, rr, rrType, rrValue string) (recordID string) {
	qcns := cns.New(cli.AKID, cli.AKEY)
	id, err := qcns.RecordCreate(
		domain,
		cns.Record{
			Name:  rr,
			Type:  strings.ToUpper(rrType),
			Value: rrValue,
		})

	if err != nil {
		logrus.Fatalln(err)
	}

	return strconv.Itoa(id)
}
