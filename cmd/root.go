package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the root command
var rootCmd = &cobra.Command{
	Use:   "dnsx",
	Short: "支持多配置的选择的 dns 命令客户端",
	Long: `支持多配置供应商的 dns 命令客户端， 通过执行不同的 profile 为域名添加解析记录。 
	默认使用 default profile。 使用 -p 切换其他 profile`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	Execute()
	// },
}

func init() {
	// rootCmd.AddCommand(rootCmd)
	rootCmd.PersistentFlags().StringVarP(&profile, "profile", "p", "default", "具体 profile")
	rootCmd.PersistentFlags().StringVarP(&conf, "conf", "c", "~/.dnsx/config.json", "配置文件路径")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
