package cmds

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/dnsx/cmd/dnsx/global"
	"github.com/tangx/dnsx/pkg/dnsxctx"
)

var configureDeleteCmd = &cobra.Command{
	Use:   "rm",
	Short: "从 Config 中删除 profile",
	Run: func(cmd *cobra.Command, args []string) {
		DeleteProfile()
	},
}

// DeleteProfile 删除
func DeleteProfile() {
	// dnsx := global.Load()
	config := dnsxctx.NewConfig(global.ConfigFile)

	var profiles []string
	for k := range config.Items {
		profiles = append(profiles, k)
	}

	fmt.Println(profiles)

	var profile string
	_ = survey.AskOne(
		&survey.Select{
			Message: "选择需要删除的 Profile",
			Options: profiles,
		},
		&profile,
	)

	fmt.Println(profile)

	confirm := false
	_ = survey.AskOne(
		&survey.Confirm{Message: fmt.Sprintf("确认删除 %s ？", profile)},
		&confirm,
	)

	if confirm {
		// config.Delete(profile)
		config.DeleteItem(profile)

		// config.Dump(global.CfgFile)
		logrus.Infof("已删除 Profile(%s)", profile)
	} else {
		logrus.Infof("用户取消删除 Profile(%s)", profile)
	}

}
