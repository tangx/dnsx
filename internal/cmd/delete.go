package cmd

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	ForceMode bool
	Confirm   bool
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "删除解析记录",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			logrus.Fatalln("参数不够: dnsx delete example.org [pattern]")
		}

		DeleteRecord(args)
	},
}

func init() {
	deleteCmd.Flags().BoolVarP(&ForceMode, "force", "f", false, "强制模式")
}

// DeleteRecord 删除解析记录
// dnsx delete example.org record , then go interactive
func DeleteRecord(args []string) {

	var domain string
	var pattern string

	domain = args[0]
	if len(args) == 1 {
		pattern = ""
	} else {
		pattern = args[1]
	}

	// 获取所有符合 record 查询条件的 Record 解析记录
	Records := dnsx.GetRecords(domain, pattern)
	if len(Records) == 0 {
		logrus.Infoln("没有找到匹配的域名解析记录")
		return
	}

	// 进入交互删除交互界面
	if ForceMode {
		logrus.Infof("强制删除模式，无需确认   请谨慎操作\n\n")
	}

	var QsRecordSelectOpts []string
	RecordsDict := make(map[string]string)

	// 组装数据
	// format := "%-15s %-20s %-8s %-20s %-10s"
	format := "%s: (%s) %10s.%s  %-5s  %-10s  %-10s"
	for _, rr := range Records {
		// 123841: www A 1.2.3.4 (enable)
		value := fmt.Sprintf(format, rr.ID, rr.Status, rr.Name, domain, rr.Type, rr.Value, rr.UpdateOn)
		QsRecordSelectOpts = append(QsRecordSelectOpts, value)
		RecordsDict[rr.ID] = value
	}

	// 准备问题
	var QsRecordAnswers []string
	QsRecordMultiSelect := &survey.MultiSelect{
		Message: "选择要删除的解析记录",
		Options: QsRecordSelectOpts,
	}

	QsMakeSure := &survey.Confirm{
		Message: "确认删除所选列表? ",
	}

	// 选择需要删除的记录对象
	_ = survey.AskOne(QsRecordMultiSelect, &QsRecordAnswers)
	if len(QsRecordAnswers) == 0 {
		logrus.Info("用户取消 或 没有选择删除对象")
		return
	}

	// 确认删除
	_ = survey.AskOne(QsMakeSure, &Confirm)
	if !Confirm {
		logrus.Infof("用户取消操作\n")
		return
	}

	// 执行删除
	for _, answer := range QsRecordAnswers {
		id := strings.Trim(strings.Split(answer, ":")[0], "")

		result := dnsx.DeleteRecord(domain, id)
		logrus.Infof("成功删除 %s", RecordsDict[result])
	}

}
