package aliyun

import (
	"fmt"
	"testing"

	"github.com/tangx/dnsx/config"
)

// https://help.aliyun.com/document_detail/29776.html
func TestDNS_InitClient(t *testing.T) {

	_, akid, akey := config.LoadDomainConfig("rockontrol.com")

	ali := AliyunDNS{
		AKID:   akid,
		AKEY:   akey,
		Domain: "rockontrol.com",
	}

	rrID := ali.add("tagnxin233333", "A", "172.31.1.31")
	fmt.Println(rrID)

}
