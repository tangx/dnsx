package main

import (
	"github.com/sirupsen/logrus"
	"github.com/tangx/dnsx/cmd/dnsx/cmds"
)

func main() {
	cmds.Execute()
}

func init() {
	logrus.SetReportCaller(true)
}
