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

	if domainConfig, ok := config.Domains[domain]; ok {
		return domainConfig.Provider, domainConfig.AKID, domainConfig.AKEY
	}

	return
}
