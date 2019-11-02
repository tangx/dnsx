package qcloud

import (
	"testing"

	"github.com/tangx/dnsx/config"
)

func TestQcloudCNS_RecordSearch(t *testing.T) {
	_, akid, akey := config.LoadDomainConfig("example.com")

	qcns := QcloudCNS{
		AKID:   akid,
		AKEY:   akey,
		Domain: "example.com",
	}

	//qcns.List()
	//qcns.Add("tangx", "A", "172.10.1.1")
	//
	//qcns.Delete("tangx")

	qcns.add("tangx2", "A", "172.10.1.1")
	qcns.Update("tangx2", "A", "192.168.0.1")
	qcns.Get("tangx2")
	qcns.Delete("tangx2")

}
