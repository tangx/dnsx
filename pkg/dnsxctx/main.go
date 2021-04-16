package dnsxctx

import (
	"fmt"
	"strings"

	"github.com/tangx/dnsx/cmd/dnsx/global"
	"github.com/tangx/dnsx/pkg/backend"
	"github.com/tangx/dnsx/pkg/backend/aliyun"
	"github.com/tangx/dnsx/pkg/backend/qcloud"
)

type Dnsx struct{}

// 根据变量获取
func NewDnsx() (backend.Provider, error) {
	// 1. 获取 flags
	// 2. 根据 profile 的值返回对应的 读取对应配置变量
	// 3. 根据变量返回 dnsx 控制

	switch strings.ToLower(global.Flags.Profile) {
	case "aliyun":
		return aliyun.NewClient("", ""), nil
	case "qcloud", "dnspod":
		return qcloud.NewClient("", ""), nil
	}

	return nil, fmt.Errorf("不支持供应商")

}
