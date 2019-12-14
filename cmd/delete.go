package cmd

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	areYouSure bool
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "删除解析记录",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			logrus.Fatalln("参数不够")
		}

		DeleteRecord(args)
	},
}

func init() {
	deleteCmd.Flags().BoolVarP(&areYouSure, "force", "f", false, "强制模式")
}

// DeleteRecord 删除解析记录
// dnsx delete example.org record , then go interactive
func DeleteRecord(args []string) {
	domain := args[0]
	pattern := args[1]

	IClient := GetClient()

	// 获取所有符合 record 查询条件的 Record 解析记录
	Records := IClient.GetRecords(domain, pattern)

	// 进入交互删除交互界面
	var promptOpts []string
	RecordsDict := make(map[string]string)

	if len(Records) == 0 {
		logrus.Infoln("没有找到匹配的域名解析记录")
		return
	}

	if areYouSure {
		logrus.Infof("强制删除模式，无需确认   请谨慎操作\n\n")
	}
	// 组装数据
	for _, rr := range Records {
		// 123841: www A 1.2.3.4 (enable)
		value := fmt.Sprintf("%s: %s %s %s (%s)", rr.ID, rr.Name, rr.Type, rr.Value, rr.Status)
		promptOpts = append(promptOpts, value)
		RecordsDict[rr.ID] = value
	}

	// 准备问题
	var Answers []string
	promtpSelect := &survey.MultiSelect{
		Message: "选择要删除的解析记录",
		Options: promptOpts,
	}

	promptSure := &survey.Confirm{
		Message: "已经确认，可以执行删除",
	}

	// 选择
	survey.AskOne(promtpSelect, &Answers)
	// 确认
	if !areYouSure {
		survey.AskOne(promptSure, &areYouSure)
	}

	for _, answer := range Answers {
		id := strings.Trim(strings.Split(answer, ":")[0], "")

		result := IClient.DeleteRecord(domain, id)
		logrus.Infof("删除 %s", RecordsDict[result])
	}

}
