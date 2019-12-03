package cmd

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"path"

	"github.com/tangx/dnsx/utils"
)

type DnsxConfig struct {
	Profile map[string]DnsxConfigItem `json:"profile,omitempty"`
}

type DnsxConfigItem struct {
	AKEY     string   `json:"AKEY,omitempty" example:"KEYxxxxx"`
	AKID     string   `json:"AKID,omitempty" example:"IDxxxxxxx"`
	Domains  []string `json:"Domains,omitempty" example:"example.org,example.com"`
	Provider string   `json:"Provider,omitempty" example:"aliyun"`
}

// LoadConfig load config
func (c Client) LoadConfig() DnsxConfigItem {

	if c.Profile == "" {
		c.Profile = "default"
	}
	if c.Config == "" || c.Config == "~/.dnsx/config.json" {
		c.Config = path.Join(utils.UserHome(), ".dnsx/config.json")
	}

	var dx DnsxConfig
	if !utils.FileExists(c.Config) {
		panic(errors.New("配置文件不存在"))
	}
	body, err := ioutil.ReadFile(c.Config)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &dx)
	if err != nil {
		panic(err)
	}

	dnsConfig, ok := dx.Profile[c.Profile]
	if ok {
		return dnsConfig
	}
	return DnsxConfigItem{}
}
