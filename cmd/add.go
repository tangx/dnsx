package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/dnsx/backend/aliyun"
	"github.com/tangx/dnsx/backend/qcloud"
	"github.com/tangx/dnsx/global"
	"strings"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "添加域名解析",
	Long:  `dnsx add [flags] domain record type value`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 4 {
			logrus.Fatalln("参数不低于4")
		}
		AddRecord(args)
	},
}

func init() {
}

// AddRecord 添加域名解析记录
func AddRecord(args []string) {
	var iClient Client
	dnsx := global.Load()
	var item global.DNSxConfigItem

	if global.Profile == "default" {
		item = dnsx.Items[dnsx.Current]
	} else {
		item = dnsx.Items[global.Profile]
	}

	// fmt.Println(item)
	switch item.Provider {
	case "aliyun":
		iClient = aliyun.Client{AKID: item.AKID, AKEY: item.AKEY}
	case "qcloud":
		iClient = qcloud.Client{AKID: item.AKID, AKEY: item.AKEY}
	default:
		logrus.Fatalf("Provider(%s) : 不支持 DNS 供应商", item.Provider)
	}

	// domain, record, rrType, value := args[0:4]
	domain := args[0]
	record := args[1]
	recordType := strings.ToUpper(args[2])
	value := args[3]
	RecordID := iClient.AddRecord(domain, record, recordType, value)

	if RecordID != "" {
		logrus.Infof("Success:[%s] %s.%s (%s) %s ", RecordID, record, domain, recordType, value)
	}
}
