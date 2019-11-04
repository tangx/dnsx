package aliyun

import (
	"testing"

	"github.com/tangx/dnsx/config"
)

// https://help.aliyun.com/document_detail/29776.html
func TestDNS_InitClient(t *testing.T) {

	_, akid, akey := config.LoadDomainConfig("example.org")

	ali := Alidns{
		AKID:   akid,
		AKEY:   akey,
		Domain: "example.org",
	}

	ali.Add("tagnxin2333", "A", "172.31.1.31")
}
