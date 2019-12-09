package qcloud

import cns "github.com/go-http/qcloud-cns"
import "github.com/tangx/dnsx/utils"

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
			Type:  rrType,
			Value: rrValue,
		})
	utils.PanicError(err)

	return string(id)
}
