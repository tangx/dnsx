package global

import (
	"os"
	"path"
)

type CommandFlags struct {
	Type    string `flag:"type" short:"t" persistent:"true" usage:"record type, ex: cname,a,txt"`
	Record  string `flag:"record" short:"r" persistent:"true" usage:"record name: ex: www"`
	Value   string `flag:"value" short:"v" persistent:"true" usage:"record value, ex: 192.168.0.1"`
	Domain  string `flag:"domain" short:"d" persistent:"true" usage:"record domain, ex: example.com"`
	Status  bool   `flag:"status" short:"s" persistent:"true" usage:"record status, ex: enable"`
	Profile string `flag:"profile" short:"p" persistent:"true" usage:"profile name"`
}

var Flags = CommandFlags{
	Type:    "A",
	Status:  true,
	Profile: "current",
}

var (
	ConfigFile = path.Join(os.Getenv("HOME"), ".dnsx", "config.yml")
)
