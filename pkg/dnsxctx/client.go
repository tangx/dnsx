package dnsxctx

import (
	"github.com/sirupsen/logrus"
	"github.com/tangx/dnsx/pkg/backend"
	"github.com/tangx/dnsx/pkg/backend/aliyun"
	"github.com/tangx/dnsx/pkg/backend/qcloud"
)

// DnsxClient for dnsx
type DnsxClient interface {
	AddRecord(domain, record, rrType, Value string) (recordID string)
	GetRecords(domain, record string) (RRs []backend.RecordItem)
	DeleteRecord(domain, recordID string) string
	SetRecordStatus(domain string, recordID string, status bool) string
}

// NewClient 根据 Provider 返回相应 DNS 客户端
func NewClient(profile string, config DnsxConfig) DnsxClient {

	if profile == "default" {
		profile = config.Current
	}

	item := config.Items[profile]

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
