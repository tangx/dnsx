package global

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Providers DNS 解析供应商
var Providers []string = []string{"aliyun", "qcloud", "dnspod"}

var (
	// CfgFile 指定配置路径
	CfgFile string
	// Profile 指定配置选项
	Profile string
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

// Load 加载配置文件
func Load() (dnsx DnsxConfig) {
	if CfgFile == "" || CfgFile == "$HOME/.dnsx/dnsx.json" {
		CfgFile = fmt.Sprintf("%s/.dnsx/dnsx.json", os.Getenv("HOME"))
	}
	data, err := ioutil.ReadFile(CfgFile)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &dnsx)
	if err != nil {
		panic(err)
	}

	return
}

// Dump 写入配置文件
func (dnsx DnsxConfig) Dump(cfgFile string) {
	f, err := os.OpenFile(cfgFile, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0700)
	if err != nil {
		panic(err)
	}

	_, _ = f.WriteString(dnsx.Marshal())
}

// New 新建 配置文件
// 这里应该使用 template 完成
func (dnsx DnsxConfig) New(cfgFile string) {}

// Marshal 格式化配置文件
func (dnsx DnsxConfig) Marshal() (s string) {

	b, err := json.MarshalIndent(dnsx, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(b)
}

// Add 增加 Profile
func (dnsx DnsxConfig) Add(profile string, item DnsxConfigItem) {}

// Delete 删除 Profile
func (dnsx *DnsxConfig) Delete(profile string) {
	delete(dnsx.Items, profile)

}
