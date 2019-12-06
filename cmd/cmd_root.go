package cmd

import (
	"github.com/spf13/cobra"
)

// global flags
var (
	cfgPath    string
	cfgProfile string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgProfile, "profile", "p", "default", "具体 profile")
	rootCmd.PersistentFlags().StringVarP(&cfgPath, "config", "c", "~/.dnsx/config.json", "配置文件路径")
}

// rootCmd represents the root command
var rootCmd = &cobra.Command{
	Use:   "dnsx",
	Short: "dns 命令客户端",
	Long:  `支持多配置供应商的 dns 命令客户端`,
}
