package cmds

import (
	"github.com/spf13/cobra"
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
	// cobrautils.BindFlags(addCmd, &global.Flags)
}

func AddRecord() {
}
