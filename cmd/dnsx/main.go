package main

import (
	"github.com/tangx/dnsx/cmd/dnsx/cmds"
)

func main() {
	cmds.Execute()
}

func init() {
	// logrus.SetReportCaller(true)
}
