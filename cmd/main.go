package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/dnsx/backend"
	"github.com/tangx/dnsx/backend/aliyun"
	"github.com/tangx/dnsx/backend/qcloud"
	"github.com/tangx/dnsx/global"
	"github.com/tangx/dnsx/version"
)

// configureCmd represents the configure command
var rootCmd = &cobra.Command{
	Use:   "dnsx",
	Short: "DNSx 配置管理 DNS 解析",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("version: " + version.Version)
	},
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(switchCmd)

	// global vars
	rootCmd.PersistentFlags().StringVarP(&global.CfgFile, "config", "c", "$HOME/.dnsx/dnsx.json", "config file")
	rootCmd.PersistentFlags().StringVarP(&global.Profile, "profile", "p", "default", "profile")
}

// Execute 执行命令
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// DnsxClient for dnsx
type DnsxClient interface {
	AddRecord(domain, record, rrType, Value string) (recordID string)
	GetRecords(domain, record string) (RRs []backend.RecordItem)
	DeleteRecord(domain, recordID string) string
	SetRecordStatus(domain string, recordID string, status bool) string
}

// getClient 根据 Provider 返回相应 DNS 客户端
func getClient() (iClient DnsxClient) {
	// var iClient Client
	dnsx := global.Load()
	var item global.DnsxConfigItem

	if global.Profile == "default" {
		item = dnsx.Items[dnsx.Current]
	} else {
		item = dnsx.Items[global.Profile]
	}

	// fmt.Println(item)
	switch item.Provider {
	case "aliyun":
		iClient = aliyun.NewClient(item.AKID, item.AKEY)
	case "qcloud":
		iClient = qcloud.NewClient(item.AKID, item.AKEY)
	default:
		logrus.Fatalf("Provider(%s) : 不支持 DNS 供应商", item.Provider)
	}

	return
}

var dnsx DnsxClient

func init() {
	dnsx = getClient()
}
