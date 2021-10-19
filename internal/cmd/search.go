package cmd

import (
	"errors"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/dnsx/pkg/backend"
)

// searchCmd represents the get command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "查询记录信息",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			logrus.Fatalf("%s", errors.New("dnsx search example.org [record]"))
		}
		GetRecords(args)
	},
}

// GetRecords 查询记录信息
func GetRecords(args []string) {

	var record string

	domain := args[0]
	if len(args) == 1 {
		record = ""
	} else {
		record = args[1]
	}

	RRs := dnsx.GetRecords(domain, record)

	// dumpByPrintf(RRs)
	tableWriter(RRs)
}

func tableWriter(RRs []backend.RecordItem) {

	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"RecordID", "Record", "Type", "Value", "Status", "Last Update Time"})

	// table.Append([]string{})

	for _, rr := range RRs {
		table.Append([]string{rr.ID, rr.Name, rr.Type, rr.Value, rr.Status, rr.UpdateOn})
	}

	table.Render()
}
