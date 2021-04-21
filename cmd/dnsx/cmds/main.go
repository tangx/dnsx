package cmds

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/cobrautils"
	"github.com/tangx/dnsx/cmd/dnsx/global"
	"github.com/tangx/dnsx/pkg/backend"
	"github.com/tangx/dnsx/pkg/dnsxctx"
	"github.com/tangx/dnsx/version"
)

var dcli backend.DnsClient

var rootCmd = &cobra.Command{
	Use:     "dnsx",
	Short:   "dnsx 一款命令行 dns 解析客户端",
	Version: fmt.Sprintf(version.Version),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		SetDefualts()
	},
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	// add subcommand
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(switchCmd)
	rootCmd.AddCommand(configureCmd)
	rootCmd.AddCommand(completionCmd)

	// add httpserver
	rootCmd.AddCommand(httpserverCmd)

	// binding flags
	cobrautils.BindFlags(rootCmd, &global.Flags)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}

}

func SetDefualts() {

	var err error
	dcli, err = dnsxctx.NewDnsxClient()
	if err != nil {
		logrus.Fatalf("Get Dnsx Client failed: %v", err)
	}

}
