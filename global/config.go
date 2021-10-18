package global

// Providers DNS 解析供应商
var Providers []string = []string{"aliyun", "qcloud", "dnspod"}

var (
	// CfgFile 指定配置路径
	CfgFile string
	// Profile 指定配置选项
	Profile string
)
