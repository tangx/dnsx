package cmds

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/cobrautils"
	"github.com/tangx/dnsx/cmd/dnsx/global"
	"github.com/tangx/dnsx/pkg/dnsxctx"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "管理配置文件",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cobrautils.BindFlags(cmd, &global.Flags)
	},
}

var configureCurrentCmd = &cobra.Command{
	Use:   "default",
	Short: "修改 current值， 设置默认生效的 profile",
	Run: func(cmd *cobra.Command, args []string) {
		SetCurrent()
	},
}

func init() {
	configureCmd.AddCommand(configureAddCmd)
	configureCmd.AddCommand(configureCurrentCmd)
	configureCmd.AddCommand(configureDeleteCmd)
}

// SetCurrent 配置默认 Profile
func SetCurrent() {

	config := dnsxctx.NewConfig(global.ConfigFile)
	if ok, err := config.SetDefault(global.Flags.Profile); !ok {
		logrus.Fatalf("%v", err)
	}
	config.Current = global.Flags.Profile
}
