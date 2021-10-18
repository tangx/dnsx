package cmd

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/dnsx/global"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "管理配置文件",
	// PersistentPreRun: func(cmd *cobra.Command, args []string) {
	// 	config = global.Load()
	// },
	// Run: func(cmd *cobra.Command, args []string) {
	// 	ConfigureMain()
	// },
}

var configureDomainsCmd = &cobra.Command{
	Use:   "domains",
	Short: "列出 Config 中的所有 Profile 中的所有域名",
	Run: func(cmd *cobra.Command, args []string) {
		Domains()
	},
}

var configureListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出 Config 中的所有 profile",
	Run: func(cmd *cobra.Command, args []string) {
		ListProfile()
	},
}

var configureAddCmd = &cobra.Command{
	Use:   "add",
	Short: "向 Config 中增加 profile",
	Run: func(cmd *cobra.Command, args []string) {
		AddProfile()
	},
}
var configureDeleteCmd = &cobra.Command{
	Use:   "rm",
	Short: "从 Config 中删除 profile",
	Run: func(cmd *cobra.Command, args []string) {
		DeleteProfile()
	},
}

var configureCurrentCmd = &cobra.Command{
	Use:   "set",
	Short: "修改 current值， 设置默认生效的 profile",
	Run: func(cmd *cobra.Command, args []string) {
		SetCurrent()
	},
}

func init() {
	configureCmd.AddCommand(configureAddCmd)
	configureCmd.AddCommand(configureCurrentCmd)
	configureCmd.AddCommand(configureDeleteCmd)
	configureCmd.AddCommand(configureListCmd)
	configureCmd.AddCommand(configureDomainsCmd)
}

// AddProfile 增加
func AddProfile() {
	config := global.Load()
	var item global.DnsxConfigItem

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
		_ = survey.Ask(qsLoginWithKey, &item)
	}

	var confirm bool = false

	_ = survey.AskOne(&survey.Confirm{
		Message: fmt.Sprintf("是否添加 %s 到配置中", global.Profile),
	}, &confirm,
	)

	if confirm {
		config.Items[global.Profile] = item
		config.Dump(global.CfgFile)
	} else {
		logrus.Infoln("用户取消添加")
	}

}

// DeleteProfile 删除
func DeleteProfile() {

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
		config.Delete(profile)

		config.Dump(global.CfgFile)
		logrus.Infof("已删除 Profile(%s)", profile)
	} else {
		logrus.Infof("用户取消删除 Profile(%s)", profile)
	}

}

// SetCurrent 配置默认 Profile
func SetCurrent() {

	if _, ok := config.Items[global.Profile]; !ok {
		logrus.Fatalf("Profile(%s) 不存在", global.Profile)
	}
	config.Current = global.Profile

	config.Dump(global.CfgFile)
}

// ListProfile 返回当前 config 中的所有 profile
func ListProfile() {
	var l []string
	for key := range config.Items {
		l = append(l, key)
	}

	fmt.Println(strings.Join(l, " "))
}

// Domains 返回当前 profile 中的所有 domain
func Domains() {

	var p string
	if global.Profile == "default" {
		p = config.Current
	} else {
		p = global.Profile
	}
	item := config.Items[p]
	fmt.Println(strings.Join(item.Domains, " "))
}
