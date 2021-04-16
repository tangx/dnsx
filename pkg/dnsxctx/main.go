package dnsxctx

import (
	"fmt"
	"strings"

	"github.com/tangx/dnsx/cmd/dnsx/global"
	"github.com/tangx/dnsx/pkg/backend"
	"github.com/tangx/dnsx/pkg/backend/aliyun"
	"github.com/tangx/dnsx/pkg/backend/qcloud"
)

// NewDnsxClient 返回一个控制器
func NewDnsxClient() (backend.DnsClient, error) {
	// 1. 获取 flags
	// 2. 根据 profile 的值返回对应的 读取对应配置变量
	config := NewConfig(global.ConfigFile)

	// 3. 获取指定的供应商
	profile := global.Flags.Profile
	if profile == "current" {
		profile = config.Current
	}

	// 4. 返回具体的 dns 供应商
	item := config.GetItem(profile)

	switch strings.ToLower(item.Provider) {
	case "aliyun":
		return aliyun.NewClient(item.AKID, item.AKEY), nil
	case "qcloud":
		return qcloud.NewClient(item.AKID, item.AKEY), nil
	}

	return nil, fmt.Errorf("不支持供应商")

}
