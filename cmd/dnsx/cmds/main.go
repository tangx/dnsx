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
	Long:    "dnsx 一款命令行 dns 解析客户端",
	Version: fmt.Sprintf("version: v%s", version.Version),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		SetDefualts()
	},
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	// add subcommand
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(deleteCmd)

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
		logrus.Fatalf("Get Dnsx Client failed: %w", err)
	}

}
