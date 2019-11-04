package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/tangx/dnsx/backend/aliyun"

	"github.com/tangx/dnsx/backend/qcloud"
	"github.com/tangx/dnsx/config"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add record, 增加域名解析",
	Long:  `增加域名解析`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("add called")
		Add(args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func Add(args []string) {
	//fmt.Println(args)
	// domain sub type value
	// ex: example.org www A 1.1.1.1
	if len(args) < 4 {
		logrus.Fatal("Add 参数不足")
		os.Exit(1)
	}

	providor, akid, akey := config.LoadDomainConfig(args[0])

	switch strings.ToLower(providor) {
	case "qcloud":
		{
			var qcns qcloud.QcloudCNS
			qcns.Add(args[0], args[1], args[2], args[3])
		}
	case "aliyun":
		{

			alidns := aliyun.Alidns{
				AKID:   akid,
				AKEY:   akey,
				Domain: args[0],
			}
			alidns.Add(args[1], args[2], args[3])
			//alidns.add
		}

	default:
		{
			fmt.Printf("DNS提供商( %s ) 不在支持列表", providor)
		}
	}
}
