package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/dnsx/global"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "管理配置文件",
	// Run: func(cmd *cobra.Command, args []string) {
	// 	ConfigureMain()
	// },
}

var configureAddCmd = &cobra.Command{
	Use:   "add",
	Short: "增加 profile",
	Run: func(cmd *cobra.Command, args []string) {
		AddProfile()
	},
}

var configureCurrentCmd = &cobra.Command{
	Use:   "current",
	Short: "增加 profile",
	Run: func(cmd *cobra.Command, args []string) {
		SetCurrent()
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
	configureCmd.AddCommand(configureAddCmd)
	configureCmd.AddCommand(configureCurrentCmd)
}

// ConfigureMain configure 子命令入口
func ConfigureMain() {

}

// AddProfile 增加
func AddProfile() {
	dnsx := global.Load()
	var item global.DNSxConfigItem

	var qsProvider = []*survey.Question{
		{
			Name: "provider",
			Prompt: &survey.Select{
				Message: "Choose a color:",
				Options: global.Providers,
			},
		},
	}

	var qsLoginWithKey = []*survey.Question{
		{
			Name: "AKID",
			Prompt: &survey.Input{
				Message: "输入 AK ID",
			},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
		{
			Name: "AKEY",
			Prompt: &survey.Password{
				Message: "输入 AK Secret: ",
			},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
	}

	err := survey.Ask(qsProvider, &item)
	if err != nil {
		panic(err)
	}

	if item.Provider == "aliyun" || item.Provider == "qcloud" {
		survey.Ask(qsLoginWithKey, &item)
	}

	var confirm bool = false

	survey.AskOne(&survey.Confirm{
		Message: fmt.Sprintf("是否添加 %s 到配置中", global.Profile),
	}, &confirm,
	)

	if confirm {
		dnsx.Items[global.Profile] = item
		dnsx.Dump(global.CfgFile)
	} else {
		logrus.Infoln("用户取消添加")
	}

}

// SetCurrent 配置默认 Profile
func SetCurrent() {
	dnsx := global.Load()

	if _, ok := dnsx.Items[global.Profile]; !ok {
		logrus.Fatal("Profile(%s) 不存在")
	}
	dnsx.Current = global.Profile

	dnsx.Dump(global.CfgFile)
}
