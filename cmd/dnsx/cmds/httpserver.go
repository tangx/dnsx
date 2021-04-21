package cmds

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/cobrautils"
	"github.com/tangx/dnsx/pkg/httpserver"
)

var httpserverCmd = &cobra.Command{
	Use:   "httpserver",
	Short: "启动 httpserver 服务器",
	Run: func(cmd *cobra.Command, args []string) {
		if err := httpserver.Run(); err != nil {
			logrus.Fatalf("Start httpserver failed: %v", err)
		}
	},
}

func init() {
	cobrautils.BindFlags(httpserverCmd, &httpserver.Flags)
}
