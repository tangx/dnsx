package cmds

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/dnsx/cmd/dnsx/global"
)

// switchCmd represents the switch command
var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "切换域名状态",
	Run: func(cmd *cobra.Command, args []string) {
		SwitchRecordStatus()
	},
}

func SwitchRecordStatus() {

	domain := global.Flags.Domain
	pattern := global.Flags.Record

	Records := dcli.GetRecords(domain, pattern)

	if len(Records) == 0 {
		logrus.Infof("没有匹配的解析记录")
		return
	}

	var QsRecordSelectOpts []string
	var QsRecordAnswers []string

	format := "%s: (%s) %10s.%s  %-5s  %-10s  %-10s"
	for _, rr := range Records {
		value := fmt.Sprintf(format, rr.ID, rr.Status, rr.Name, domain, rr.Type, rr.Value, rr.UpdateOn)
		QsRecordSelectOpts = append(QsRecordSelectOpts, value)
	}
	QsRecordMultiSelect := &survey.MultiSelect{
		Message: "选择域名",
		Options: QsRecordSelectOpts,
	}

	survey.AskOne(QsRecordMultiSelect, &QsRecordAnswers)

	if len(QsRecordAnswers) == 0 {
		logrus.Infoln("用户取消或无记录选择")
		return
	}

	// 选择状态
	QsRecordStatus := &survey.Select{
		Message: "选择域名状态",
		Options: []string{"enable", "disable"},
	}
	var QsRecordStatusAnswer string
	survey.AskOne(QsRecordStatus, &QsRecordStatusAnswer)

	var status bool
	if QsRecordStatusAnswer == "enable" {
		status = true
	} else {
		status = false
	}

	for _, answer := range QsRecordAnswers {
		rrID := strings.Split(answer, ":")[0]
		result := dcli.SetRecordStatus(domain, rrID, status)
		fmt.Printf("解析记录 %s 状态被设置为 %s\n", result, QsRecordStatusAnswer)
	}
}
