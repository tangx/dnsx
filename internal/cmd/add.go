package cmd

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
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
			Usage := `Usage: dnsx add type domain record value
	ex: dnsx add A example.org www 1.1.1.1`
			logrus.Fatalln(Usage)
			// logrus.Fatalln("参数不低于4")
		}

		AddRecord(args)
	},
}

func init() {
}

// AddRecord 添加域名解析记录
func AddRecord(args []string) {

	// domain, record, rrType, value := args[0:4]

	recordType := strings.ToUpper(args[0])
	domain := args[1]
	record := args[2]
	value := args[3]

	var QsRecordSelectOpts []string
	var QsRecordAnswers []string
	RecordsDict := make(map[string]string)

	Records := dnsx.GetRecords(domain, record)
	var confirm bool
	if len(Records) != 0 {
		for _, rr := range Records {
			if rr.Name == record && rr.Status == "ENABLE" {
				//准备数据
				format := "%s: (%s) %10s.%s  %-5s  %-10s  %-10s"
				value := fmt.Sprintf(format, rr.ID, rr.Status, rr.Name, domain, rr.Type, rr.Value, rr.UpdateOn)
				QsRecordSelectOpts = append(QsRecordSelectOpts, value)
				RecordsDict[rr.ID] = value
			}
		}
	}
	if len(QsRecordSelectOpts) != 0 {

		message := fmt.Sprintf("发现域名: [%s.%s], 已存在相关解析! 请根据提示继续执行.\ndelete :  删除当前已存在域名解析后再进行添加.\ndisable: 将当前已存在域名解析状态置为disable后再进行添加.\nappend :  继续添加同名的域名解析.\n", record, domain)
		// 选择执行动作
		QsRecordAction := &survey.Select{
			Message: message,
			Options: []string{"delete", "disable", "append"},
		}
		var QsRecordActionAnswer string
		_ = survey.AskOne(QsRecordAction, &QsRecordActionAnswer)

		if len(QsRecordActionAnswer) == 0 {
			logrus.Infoln("用户取消或无记录选择")
			return
		}

		switch QsRecordActionAnswer {
		case "delete":
			//选择域名
			QsRecordMultiSelect := &survey.MultiSelect{
				Message: "选择域名",
				Options: QsRecordSelectOpts,
			}

			_ = survey.AskOne(QsRecordMultiSelect, &QsRecordAnswers)

			if len(QsRecordAnswers) == 0 {
				logrus.Infoln("用户取消或无记录选择")
				return
			}
			QsDelMakeSure := &survey.Confirm{
				Message: "确认删除所选列表? ",
			}
			// 确认删除
			_ = survey.AskOne(QsDelMakeSure, &confirm)
			if !confirm {
				logrus.Infof("用户取消操作\n")
				return
			}
			// 执行删除
			for _, answer := range QsRecordAnswers {
				id := strings.Trim(strings.Split(answer, ":")[0], "")

				result := dnsx.DeleteRecord(domain, id)
				logrus.Infof("成功删除 %s", RecordsDict[result])
			}
		case "disable":
			QsRecordMultiSelect := &survey.MultiSelect{
				Message: "选择域名",
				Options: QsRecordSelectOpts,
			}

			_ = survey.AskOne(QsRecordMultiSelect, &QsRecordAnswers)

			if len(QsRecordAnswers) == 0 {
				logrus.Infoln("用户取消或无记录选择")
				return
			}
			QsDisableMakeSure := &survey.Confirm{
				Message: "确认Disable所选列表? ",
			}
			// 确认删除
			_ = survey.AskOne(QsDisableMakeSure, &confirm)
			if !confirm {
				logrus.Infof("用户取消操作\n")
				return
			}

			status := false

			for _, answer := range QsRecordAnswers {
				rrID := strings.Split(answer, ":")[0]
				result := dnsx.SetRecordStatus(domain, rrID, status)
				fmt.Printf("解析记录 %s 状态被设置为 Disable\n", result)
			}
		case "append":
			fmt.Printf("继续追加一条同名解析记录！!\n")
		}
	}
	RecordID := dnsx.AddRecord(domain, record, recordType, value)

	if RecordID != "" {
		logrus.Infof("成功添加:[%s] %s.%s (%s) %s ", RecordID, record, domain, recordType, value)
	}
}
