package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

// Config DNSX配置文件
type Config struct {
	Domains map[string]DomainConfig `json:"domains,omitempty"`
}

// DomainConfig 域名信息配置文件
type DomainConfig struct {
	AKEY     string `json:"AKEY,omitempty"`
	AKID     string `json:"AKID,omitempty"`
	Provider string `json:"Provider,omitempty"`
}

//const cfgfile = "~/.dnsx/dnsx.json"

// Configfile return config file path
func Configfile() string {
	home, _ := os.UserHomeDir()
	return path.Join(home, string(os.PathSeparator), ".dnsx/dnsx.json")
}

// InitCfg init config if not exist
func (c Config) InitCfg() {

}

// LoadConfig Load config file
func LoadConfig() (config Config) {

	cfgfile := Configfile()

	fb, err := ioutil.ReadFile(cfgfile)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fb, &config)
	if err != nil {
		panic(err)
	}

	return
}

// LoadDomainConfig return domain config info
func LoadDomainConfig(domain string) (provider, akid, akey string) {

	config := LoadConfig()

	// 可以通过 ok 判断确认某个键是否在 list(slice) 或 map 中。
	//   如果存在，则返回 value, true
	//   如果不存在， 则返回  0值, false
	//
	// 注意1: struct 不能使用 ok 判断，
	// 注意2: 指针值使用键值的时候，需要加圆括号
	//    https://stackoverflow.com/questions/25290956/go-update-slice-iterating-error-does-not-support-indexing
	//    https://flaviocopes.com/golang-does-not-support-indexing/
	//    invalid operation: members[0] (type *Members does not support indexing)
	if domainConfig, ok := config.Domains[domain]; ok {
		return domainConfig.Provider, domainConfig.AKID, domainConfig.AKEY
	}

	return
}
