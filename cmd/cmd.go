package cmd

import (
	"log"

	"github.com/tangx/dnsx/backend/aliyun"
	"github.com/tangx/dnsx/backend/qcloud"
)

var (
	conf    string
	profile string

	dnsx Dnsx
)

// Client for dnsx
type Client struct {
	Config  string
	Profile string
}

// Dnsx actions
type Dnsx interface {
	Add(domain, rr, rrType, rrValue string) string
}

// InitConfig to run
func InitConfig() {
	c := Client{
		Config:  conf,
		Profile: profile,
	}

	cfg := c.LoadConfig()

	switch cfg.Provider {
	case "aliyun":
		{
			dnsx = aliyun.AliyunDNS{
				AKID: cfg.AKID,
				AKEY: cfg.AKEY,
			}
		}
	case "qcloud":
		{
			dnsx = qcloud.QcloudCNS{
				AKID: cfg.AKID,
				AKEY: cfg.AKEY,
			}
		}
	default:
		{
			log.Fatal("没有或不支持的 Provider")
		}
	}
}
