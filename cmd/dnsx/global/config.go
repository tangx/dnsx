package global

import (
	"os"
	"path"
)

type CommandFlags struct {
	Type    string `flag:"type" short:"t" usage:"record type, ex: cname,a,txt"`
	Record  string `flag:"record" short:"v" usage:"record name: ex: www"`
	Value   string `flag:"value" short:"v" usage:"record value, ex: 192.168.0.1"`
	Domain  string `flag:"domain" short:"d" usage:"record domain, ex: example.com"`
	Status  bool   `flag:"status" short:"s" usage:"record status, ex: enable"`
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
