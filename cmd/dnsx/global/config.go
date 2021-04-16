package global

import (
	"os"
	"path"
)

type CommandFlags struct {
	Type    string `flag:"type" shorthand:"t" persistent:"true" usage:"record type, ex: cname,a,txt"`
	Record  string `flag:"record" shorthand:"r" persistent:"true" usage:"record name: ex: www"`
	Value   string `flag:"value" shorthand:"v" persistent:"true" usage:"record value, ex: 192.168.0.1"`
	Domain  string `flag:"domain" shorthand:"d" persistent:"true" usage:"record domain, ex: example.com"`
	Status  bool   `flag:"status" shorthand:"s" persistent:"true" usage:"record status, ex: enable"`
	Profile string `flag:"profile" shorthand:"p" persistent:"true" usage:"profile name"`
}

var Flags = CommandFlags{
	Type:    "A",
	Status:  true,
	Profile: "current",
}

var (
	ConfigFile = path.Join(os.Getenv("HOME"), ".dnsx", "config.yml")
)
