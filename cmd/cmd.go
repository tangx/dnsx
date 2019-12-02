package cmd

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/tangx/dnsx/utils"
)

type Dnsx struct {
	Profile map[string]DnsxConfig `json:"profile,omitempty"`
}

type DnsxConfig struct {
	AKEY     string   `json:"AKEY,omitempty" example:"KEYxxxxx"`
	AKID     string   `json:"AKID,omitempty" example:"IDxxxxxxx"`
	Domains  []string `json:"Domains,omitempty" example:"example.org,example.com"`
	Provider string   `json:"Provider,omitempty" example:"aliyun"`
}

// Client for dnsx
type Client struct {
	Config string
}

func (c Client) LoadConfig(profile string) DnsxConfig {
	var dx Dnsx
	if !utils.FileExists(c.Config) {
		panic(errors.New("config is not exists"))
	}
	body, err := ioutil.ReadFile(c.Config)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &dx)
	if err != nil {
		panic(err)
	}

	dnsConfig, ok := dx.Profile[profile]
	if ok {
		return dnsConfig
	}
	return DnsxConfig{}
}
