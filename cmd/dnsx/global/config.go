package global

type CommandFlags struct {
	Type    string `flag:"type" shorthand:"t" usage:"record type, ex: cname,a,txt"`
	Name    string `flag:"name" shorthand:"n" usage:"record name: ex: www"`
	Value   string `flag:"value" shorthand:"v" usage:"record value, ex: 192.168.0.1"`
	Domain  string `flag:"domain" shorthand:"d" usage:"record domain, ex: example.com"`
	Status  bool   `flag:"status" shofthand:"s" usage:"record status, ex: enable"`
	Profile string `flag:"profile" shorthand:"p" persistent:"true" usage:"profile name"`
}

var Flags = CommandFlags{
	Type:    "A",
	Status:  true,
	Profile: "current",
}
