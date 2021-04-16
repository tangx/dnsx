package cmds

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tangx/cobrautils"
	"github.com/tangx/dnsx/cmd/dnsx/global"
	"github.com/tangx/dnsx/version"
)

var rootCmd = &cobra.Command{
	Use:  "dnsx",
	Long: fmt.Sprintf("version: v%s", version.Version),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	cobrautils.BindFlags(rootCmd, &global.Flags)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
