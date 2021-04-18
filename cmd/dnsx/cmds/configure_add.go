package cmds

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/dnsx/cmd/dnsx/global"
	"github.com/tangx/dnsx/pkg/backend"
	"github.com/tangx/dnsx/pkg/dnsxctx"
)

var configureAddCmd = &cobra.Command{
	Use:   "add",
	Short: "向 Config 中增加 profile",

	Run: func(cmd *cobra.Command, args []string) {
		AddProfile()
	},
}

// AddProfile 增加
func AddProfile() {
	// dnsx := global.Load()
	// var item global.DNSxConfigItem
	var item dnsxctx.DnsxConfigItem

	var qsProvider = []*survey.Question{
		{
			Name: "provider",
			Prompt: &survey.Select{
				Message: "Choose a color:",
				Options: backend.Providers,
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
		_ = survey.Ask(qsLoginWithKey, &item)
	}

	var confirm bool = false

	_ = survey.AskOne(&survey.Confirm{
		Message: fmt.Sprintf("是否添加 %s 到配置中", global.Flags.Profile),
	}, &confirm)

	// fmt.Println(item)
	if confirm {
		config := dnsxctx.NewConfig(global.ConfigFile)
		config.AddItem(global.Flags.Profile, item)
	} else {
		logrus.Infoln("用户取消添加")
	}

}
