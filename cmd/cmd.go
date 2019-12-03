package cmd

import (
	"flag"
	"log"
	"strings"

	"github.com/tangx/dnsx/backend/aliyun"
	"github.com/tangx/dnsx/backend/qcloud"
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

var profile = flag.String("profile", "default", "操作的 profile 的名称")
var conf = flag.String("conf", "~/.dnsx/config.json", "配置文件")

// Main to run
func Main() {
	flag.Parse()
	c := Client{
		Config:  *conf,
		Profile: *profile,
	}

	var dnsx Dnsx
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

	args := flag.Args()
	if len(args) < 1 {

		usage := `-profile default set example.com www a 1.2.3.4`
		log.Fatalln(usage)
	}

	switch strings.ToLower(args[0]) {
	case "add":
		{
			dnsx.Add(args[1], args[2], args[3], args[4])
		}
	}
}
