package dnsxctx

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/tangx/dnsx/cmd/dnsx/global"
	"gopkg.in/yaml.v2"
)

type DnsxConfigItem struct {
	AKID     string   `json:"akid,omitempty" yaml:"akid,omitempty"`
	AKEY     string   `json:"akey,omitempty" yaml:"akey,omitempty"`
	Provider string   `json:"provider,omitempty" yaml:"provider,omitempty"`
	Domains  []string `json:"domains,omitempty" yaml:"domains,omitempty"`
	Comment  string   `json:"comment,omitempty" yaml:"comment,omitempty" usage:"备注"`
}

type DnsxConfig struct {
	Current string                    `json:"current,omitempty" yaml:"current,omitempty" usage:"默认配置"`
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

func (dc *DnsxConfig) GetItem(name string) DnsxConfigItem {
	return dc.Items[name]
}

func (dc *DnsxConfig) AddItem(name string, item DnsxConfigItem) {
	dc.Items[name] = item
	dc.DumpYaml()
}

func (dc *DnsxConfig) DeleteItem(name string) {
	_, ok := dc.Items[name]
	if ok {
		delete(dc.Items, name)
	}

	dc.DumpYaml()
}

func (dc *DnsxConfig) DumpYaml() {
	data, err := yaml.Marshal(dc)
	if err != nil {
		logrus.Fatalf("marshal yaml config failed, %v", err)
	}

	f, err := os.OpenFile(global.ConfigFile, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		logrus.Fatalf("open config file failed, %v", err)
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		logrus.Fatalf("dump config to file failed, %v", err)
	}

}
