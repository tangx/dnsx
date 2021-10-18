package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/tangx/dnsx/internal/global"
	"github.com/tangx/dnsx/pkg/dnsxctx"
	"github.com/tangx/dnsx/version"
)

var dnsx dnsxctx.DnsxClient
var config dnsxctx.DnsxConfig

// configureCmd represents the configure command
var rootCmd = &cobra.Command{
	Use: "dnsx",
	Short: fmt.Sprintf(`DNSx 配置管理 DNS 解析
	version: %s`, version.Version),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		convertConfigPath()

		config = dnsxctx.NewConfig(global.CfgFile)
		dnsx = dnsxctx.NewClient(global.Profile, config)
	},
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(switchCmd)

	// global vars
	rootCmd.PersistentFlags().StringVarP(&global.CfgFile, "config", "c", "$HOME/.dnsx/dnsx.json", "config file")
	rootCmd.PersistentFlags().StringVarP(&global.Profile, "profile", "p", "default", "profile")
}

// Execute 执行命令
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func convertConfigPath() {
	if global.CfgFile == "" || global.CfgFile == "$HOME/.dnsx/dnsx.json" {
		global.CfgFile = filepath.Join(os.Getenv("HOME"), ".dnsx/dnsx.json")
	}
}
