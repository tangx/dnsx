package aliyun

import (
	"fmt"
	"testing"
	"time"

	"github.com/tangx/dnsx/config"
)

// https://help.aliyun.com/document_detail/29776.html
func TestDNS_InitClient(t *testing.T) {

	domain := "example.org"
	_, akid, akey := config.LoadDomainConfig(domain)

	ali := AliyunDNS{
		AKID:   akid,
		AKEY:   akey,
		Domain: domain,
	}

	rrID := ali.Add("tagnxin213", "A", "172.31.1.31")
	fmt.Println(rrID)

	time.Sleep(1 * time.Second)

	ali.Delete("tangx")
}
