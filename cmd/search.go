package cmd

import (
	"fmt"
	"os"
	"text/template"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangx/dnsx/backend"
)

// searchCmd represents the get command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "查询记录信息",

	Run: func(cmd *cobra.Command, args []string) {
		// dnsx get example.com record
		if len(args) < 2 {
			logrus.Fatalln("参数不低于 2")
		}
		GetRecords(args)
	},
}

// GetRecords 查询记录信息
func GetRecords(args []string) {
	IClient := GetClient()

	domain := args[0]
	record := args[1]

	RRs := IClient.GetRecords(domain, record)

	dumpByPrintf(RRs)
}

func dumpByTemplate(RRs []backend.RecordItem) {

	tepl := `{{range .}}{{.ID}} {{.Name}} {{.Type}} {{.Value}} {{.Status}}
{{end}}`

	tmpl, _ := template.New("records").Parse(tepl)
	tmpl.Execute(os.Stdout, RRs)

}

func dumpByPrintf(RRs []backend.RecordItem) {

	// 彩色输出
	// https://blog.csdn.net/w616589292/article/details/51078787
	// colorFormat := "%c[1;31;41m%-20s %-20s %-8s %-20s %-10s%c[0m\n"

	format := "%-15s %-20s %-8s %-20s %-10s\n"

	fmt.Println("")
	fmt.Printf(format, "RecordID", "Record", "Type", "Value", "Status")
	fmt.Println("")
	for _, rr := range RRs {
		fmt.Printf(format, rr.ID, rr.Name, rr.Type, rr.Value, rr.Status)
	}
	fmt.Println("")
}
