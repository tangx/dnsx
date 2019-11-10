package aliyun

import (
	"fmt"
	"testing"

	"github.com/tangx/dnsx/config"
)

// https://help.aliyun.com/document_detail/29776.html
func TestDNS_InitClient(t *testing.T) {

	_, akid, akey := config.LoadDomainConfig("example.org")

	ali := AliyunDNS{
		AKID:   akid,
		AKEY:   akey,
		Domain: "example.org",
	}

	rrID := ali.Add("tagnxin213", "A", "172.31.1.31")
	fmt.Println(rrID)

}
