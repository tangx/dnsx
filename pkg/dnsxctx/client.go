package dnsxctx

import (
	"github.com/sirupsen/logrus"
	"github.com/tangx/dnsx/backend"
	"github.com/tangx/dnsx/backend/aliyun"
	"github.com/tangx/dnsx/backend/qcloud"
	"github.com/tangx/dnsx/global"
)

// DnsxClient for dnsx
type DnsxClient interface {
	AddRecord(domain, record, rrType, Value string) (recordID string)
	GetRecords(domain, record string) (RRs []backend.RecordItem)
	DeleteRecord(domain, recordID string) string
	SetRecordStatus(domain string, recordID string, status bool) string
}

// NewClient 根据 Provider 返回相应 DNS 客户端
func NewClient(config DnsxConfig) DnsxClient {

	item := config.Items[config.Current]
	if global.Profile != "default" {
		item = config.Items[global.Profile]
	}

	// fmt.Println(item)
	switch item.Provider {
	case "aliyun":
		return aliyun.NewClient(item.AKID, item.AKEY)
	case "qcloud":
		return qcloud.NewClient(item.AKID, item.AKEY)
	default:
		logrus.Fatalf("Provider(%s) : 不支持 DNS 供应商", item.Provider)
	}

	return nil
}
