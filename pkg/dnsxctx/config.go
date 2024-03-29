package dnsxctx

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// DnsxConfig 配置文件信息
type DnsxConfig struct {
	// Current 当前默认的 Profile 选项
	Current string                    `json:"current"`
	Items   map[string]DnsxConfigItem `json:"items"`
}

// DnsxConfigItem 具体变量信息
type DnsxConfigItem struct {
	AKID     string   `json:"akid,omitempty"`
	AKEY     string   `json:"akey,omitempty"`
	Provider string   `json:"provider,omitempty"`
	Domains  []string `json:"domains,omitempty"`

	// Token 为 DNSPod 的账户信息
	Token string `json:"token,omitempty"`
	// Token 为 DNSPod 的地域信息
	// CN: dnspod.cn
	// Global: dnpod.com
	Region string `json:"region,omitempty"`
}

// NewConfig 加载配置文件
func NewConfig(cfg string) (config DnsxConfig) {

	data, err := ioutil.ReadFile(cfg)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	return
}

// Dump 写入配置文件
func (config DnsxConfig) Dump(cfgFile string) {
	f, err := os.OpenFile(cfgFile, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0700)
	if err != nil {
		panic(err)
	}

	_, _ = f.WriteString(config.Marshal())
}

// New 新建 配置文件
// 这里应该使用 template 完成
func (config DnsxConfig) New(cfgFile string) {}

// Marshal 格式化配置文件
func (config DnsxConfig) Marshal() (s string) {

	b, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(b)
}

// Add 增加 Profile
func (config DnsxConfig) Add(profile string, item DnsxConfigItem) {}

// Delete 删除 Profile
func (config *DnsxConfig) Delete(profile string) {
	delete(config.Items, profile)
}
