package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tangx/dnsx/global"
)

// configureCmd represents the configure command
var rootCmd = &cobra.Command{
	Use:   "dnsx",
	Short: "DNSx 配置管理 DNS 解析",
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("configure called")
	// },
}

func init() {
	rootCmd.AddCommand(configureCmd)

	rootCmd.PersistentFlags().StringVarP(&global.CfgFile, "config", "c", "$HOME/.dnsx/dnsx.json", "config file")
	rootCmd.PersistentFlags().StringVarP(&global.Profile, "profile", "p", "defualt", "profile")
}

// Execute 执行命令
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
