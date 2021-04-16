package cmds

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/dnsx/cmd/dnsx/global"
	"github.com/tangx/dnsx/pkg/backend/http/response"
)

var searchCmd = &cobra.Command{
	Use:  "search",
	Long: `Add Dns Record`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
	Run: func(cmd *cobra.Command, args []string) {
		Search()
	},
}

func init() {
	// cobrautils.BindFlags(searchCmd, &global.Flags)
}

func Search() {

	flag := global.Flags

	domain := flag.Domain
	record := flag.Record

	if len(domain) == 0 {
		logrus.Errorf("error domain or record: domain ->(%s); record ->(%s)", domain, record)
		os.Exit(1)
	}

	rrs := dcli.GetRecords(domain, record)
	output(rrs)
}

func output(rrs []response.RecordItem) {

	// 彩色输出
	// https://blog.csdn.net/w616589292/article/details/51078787
	// colorFormat := "%c[1;31;41m%-20s %-20s %-8s %-20s %-10s%c[0m\n"

	format := "%-20s %-35s %-8s %-20s %-10s %-10s\n"

	fmt.Println("")
	fmt.Printf(format, "RecordID", "Record", "Type", "Value", "Status", "Last Update Time")
	fmt.Println("")
	for _, rr := range rrs {
		fmt.Printf(format, rr.ID, rr.Name, rr.Type, rr.Value, rr.Status, rr.UpdateOn)
	}
	fmt.Println("")
}
