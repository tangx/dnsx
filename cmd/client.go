package cmd

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/tangx/dnsx/backend/aliyun"
	"github.com/tangx/dnsx/backend/qcloud"
	"github.com/tangx/dnsx/utils"
)

type DNSxConfig struct {
	Profile map[string]DNSxConfigItem `json:"profile"`
}

type DNSxConfigItem struct {
	AKEY     string   `json:"AKEY,omitempty" example:"KEYxxxxx"`
	AKID     string   `json:"AKID,omitempty" example:"IDxxxxxxx"`
	Domains  []string `json:"Domains,omitempty" example:"example.org,example.com"`
	Provider string   `json:"Provider,omitempty" example:"aliyun"`
}

// Client for dnsx Client
type Client struct {
	Config  string
	Profile string
}

// iDNSx interface
type iDNSx interface {
	Add(domain, rr, rrType, rrValue string) string
}

// LoadConfig load config
func (c Client) LoadConfig() DNSxConfigItem {

	var dx DNSxConfig
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
	return DNSxConfigItem{}
}

func InitConfig() (dnsx iDNSx) {
	cfg := LoadConfig().Profile[cfgProfile]

	switch cfg.Provider {
	case "aliyun":
		dnsx = aliyun.AliyunDNS{AKID: cfg.AKID, AKEY: cfg.AKEY}
	case "qcloud":
		dnsx = qcloud.QcloudCNS{AKID: cfg.AKID, AKEY: cfg.AKEY}
	default:
		log.Fatal("没有或不支持的 Provider")
	}
	return
}

// ParseFlags to transfer flags
func ParseFlags() {
	if cfgProfile == "" {
		cfgProfile = "default"
	}
	if cfgPath == "" || cfgPath == "~/.dnsx/config.json" {
		cfgPath = path.Join(utils.UserHome(), ".dnsx/config.json")
	}
}

// Execute for main()
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func (dnsx DNSxConfig) Marshal() string {
	b, _ := json.MarshalIndent(dnsx, "", "  ")
	return string(b)
}

// DeleteProfile from cfgPath
func (dnsx DNSxConfig) DeleteProfile(profile string) {
	if _, ok := dnsx.Profile[profile]; ok {
		delete(dnsx.Profile, profile)
		dnsx.DumpConfig()
	}
}

// DumpConfig will store DNSxConfig to cfgPath
func (dnsx DNSxConfig) DumpConfig() {
	// f, err := os.OpenFile(cfgPath, os.O_CREATE|os.O_RDWR, 0644)
	f, err := os.OpenFile(cfgPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := dnsx.Marshal()
	f.WriteString(s)
}

// LoadConfig from cfgPath
func LoadConfig() (dnsx DNSxConfig) {
	ParseFlags()

	if !utils.FileExists(cfgPath) {
		dnsx.New(cfgPath)
	}
	cfgByte, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(cfgByte, &dnsx)
	if err != nil {
		return DNSxConfig{map[string]DNSxConfigItem{}}
	}

	return
}

// New a dnsx config
func (dnsx DNSxConfig) New(path string) {
	dnsx.DumpConfig()
}
