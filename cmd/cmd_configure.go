package cmd

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "配置管理",
	// Run: func(cmd *cobra.Command, args []string) {
	// 	AddProfile()
	// },
}

// configure subcommand add
var configureAddCmd = &cobra.Command{
	Use:   "add",
	Short: "添加管理",
	Run: func(cmd *cobra.Command, args []string) {
		AddProfile()
	},
}

// configure subcommand delete
var configureDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "删除配置",
	Run: func(cmd *cobra.Command, args []string) {
		DeleteProfile()
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
	configureCmd.AddCommand(configureAddCmd)
	configureCmd.AddCommand(configureDeleteCmd)
}

var (
	providers = []string{"qcloud", "aliyun", "dnspod"}
)

// QsProvider to select a dns provider
var QsProvider = []*survey.Question{
	{
		Name: "Provider",
		Prompt: &survey.Select{
			Message: "请指定供应商: ",
			Options: providers,
		},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
}

// QsLoginWithAccKey for aliyun, qcloud dns provider
var QsLoginWithAccKey = []*survey.Question{
	{
		Name:      "AKID",
		Prompt:    &survey.Input{Message: "AKID: "},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name:      "AKEY",
		Prompt:    &survey.Password{Message: "AKEY: "},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
}

// QsLoginWithToken for dnspod
var QsLoginWithToken = []*survey.Question{}

// AddProfile into cfgPath , rewrite if cfgProfile exists
func AddProfile() {
	var item DNSxConfigItem

	err := survey.Ask(QsProvider, &item)
	if err != nil {
		panic(err)
	}

	if item.Provider == "aliyun" || item.Provider == "qcloud" {
		err := survey.Ask(QsLoginWithAccKey, &item)
		if err != nil {
			panic(err)
		}
	}

	if item.Provider == "dnspod" {
		err := survey.Ask(QsLoginWithToken, &item)
		if err != nil {
			panic(err)
		}
	}

	dnsx := LoadConfig()
	dnsx.Profile[cfgProfile] = item

	dnsx.DumpConfig()
}

// DeleteProfile from cfgPath
func DeleteProfile() {
	dnsx := LoadConfig()

	if _, ok := dnsx.Profile[cfgProfile]; ok {
		delete(dnsx.Profile, cfgProfile)
	} else {
		log.Printf("找到不到Profile: %s", cfgProfile)
		os.Exit(0)
	}

	// QsComfirm to make confirm
	var confirm bool = false
	QsConfirm := &survey.Confirm{
		Message: fmt.Sprintf("存在 %s, 是否删除？", cfgProfile),
	}

	survey.AskOne(QsConfirm, &confirm)
	if confirm {
		dnsx.DumpConfig()
		log.Printf("删除 %s 成功", cfgProfile)
	} else {
		log.Println("取消删除")
	}
}
