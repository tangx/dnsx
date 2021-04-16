package cmds

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/dnsx/cmd/dnsx/global"
)

var addCmd = &cobra.Command{
	Use:  "add",
	Long: `Add Dns Record`,
	Run: func(cmd *cobra.Command, args []string) {
		AddRecord()
	},
}

// AddRecord 添加域名解析记录
func AddRecord() {

	flag := global.Flags
	recordType := strings.ToUpper(flag.Type)
	domain := flag.Domain
	value := flag.Value
	record := flag.Record

	if domain == "" || value == "" || record == "" || recordType == "" {
		logrus.Error("params error")
		os.Exit(1)
	}

	RecordID := dcli.AddRecord(domain, record, recordType, value)

	if RecordID != "" {
		logrus.Infof("Success:[%s] %s.%s (%s) %s ", RecordID, record, domain, recordType, value)
	}
}
