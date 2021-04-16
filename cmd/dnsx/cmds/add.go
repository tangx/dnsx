package cmds

import (
	"github.com/spf13/cobra"
	"github.com/tangx/cobrautils"
	"github.com/tangx/dnsx/cmd/dnsx/global"
)

var addCmd = &cobra.Command{
	Use:  "add",
	Long: `Add Dns Record`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	cobrautils.BindFlags(addCmd, &global.Flags)
}

func AddRecord() {
}
