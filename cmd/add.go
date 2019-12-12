package cmd

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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

	IClient := GetClient()

	// domain, record, rrType, value := args[0:4]
	domain := args[0]
	record := args[1]
	recordType := strings.ToUpper(args[2])
	value := args[3]
	RecordID := IClient.AddRecord(domain, record, recordType, value)

	if RecordID != "" {
		logrus.Infof("Success:[%s] %s.%s (%s) %s ", RecordID, record, domain, recordType, value)
	}
}
