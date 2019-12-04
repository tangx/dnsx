package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "添加域名解析记录",
	Run: func(cmd *cobra.Command, args []string) {
		addRecord(args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addRecord(args []string) {

	// fmt.Println(args)
	if len(args) != 4 {
		addHelp()
	}

	dnsx.Add(args[0], args[1], args[2], args[3])
}

func addHelp() {
	s := `
Usage: 
  dnsx add [flags] domain record type value
Example: 
  dnsx add example.com www A 1.1.3.4 
  dnsx add example.com www cname www.baidu.com -c /path/to/config.json -p profile1
More:
  use "dnsx add -h" get more 
`
	fmt.Println(s)
	os.Exit(1)
}
