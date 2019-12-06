package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/tangx/dnsx/utils"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "配置管理",
	Run: func(cmd *cobra.Command, args []string) {
		SetProfile()
	},
}

// configure subcommand add
var configureAddCmd = &cobra.Command{
	Use:   "add",
	Short: "添加管理",
	Run: func(cmd *cobra.Command, args []string) {
		SetProfile()
	},
}

// configure subcommand delete
var configureDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "删除配置",
	Run: func(cmd *cobra.Command, args []string) {
		SetProfile()
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
	configureCmd.AddCommand(configureAddCmd)
	configureCmd.AddCommand(configureDeleteCmd)
}

var (
	providers = []string{"qcloud", "aliyun"}
)

var qs = []*survey.Question{
	{
		Name: "Provider",
		Prompt: &survey.Select{
			Message: "请指定供应商: ",
			Options: providers,
		},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
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

// SetProfile to ask a profile qustion
func SetProfile() {
	ParseFlags()

	newProfile := DnsxConfigItem{}
	err := survey.Ask(qs, &newProfile)
	utils.IsError(err)

	body, err := ioutil.ReadFile(cfgPath)
	utils.IsError(err)

	var dx DNSxConfig
	json.Unmarshal(body, &dx)

	dx.Profile[cfgProfile] = newProfile

	dxbyte, _ := json.MarshalIndent(dx, "", "  ")
	fmt.Println(string(dxbyte))

	dx.DumpConfig()

}
