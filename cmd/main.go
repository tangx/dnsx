package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dnsx",
	Short: "dnsx 用于配置 dns 解析",
	Long: `dnsx 用于配置 dns 解析
			目前支持 1. qcloud cns
					2. aliyun dns`,

	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
