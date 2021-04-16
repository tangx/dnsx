package dnsxctx

import (
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type DnsxConfigItem struct {
	AKID     string   `json:"akid,omitempty" yaml:"akid,omitempty"`
	AKEY     string   `json:"akey,omitempty" yaml:"akey,omitempty"`
	Provider string   `json:"provider,omitempty" yaml:"provider,omitempty"`
	Domains  []string `json:"domains,omitempty" yaml:"domains,omitempty"`
}

type DnsxConfig struct {
	Current string                    `json:"current,omitempty" yaml:"current,omitempty"`
	Items   map[string]DnsxConfigItem `json:"items,omitempty" yaml:"items,omitempty"`
}

func NewConfig(filename string) DnsxConfig {
	dc := DnsxConfig{}
	data, err := os.ReadFile(filename)
	if err != nil {
		logrus.Fatalf("load config file failed: %v", err)
	}

	// err = json.Unmarshal(data, &dc)
	err = yaml.Unmarshal(data, &dc)
	if err != nil {
		logrus.Fatalf("unmarshal config by json failed: %v", err)
	}
	return dc
}

func (dc DnsxConfig) GetItem(name string) DnsxConfigItem {
	return dc.Items[name]
}
