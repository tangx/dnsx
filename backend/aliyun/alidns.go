package aliyun

import (
	"encoding/json"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

// Alidns is cns
type Alidns struct {
	AKID   string
	AKEY   string
	Domain string
}

type ErrorMessage struct {
	Code      string `json:"Code,omitempty"`
	HostID    string `json:"HostId,omitempty"`
	Message   string `json:"Message,omitempty"`
	RequestID string `json:"RequestId,omitempty"`
}

func getMessage(origin []byte) (result string) {

	var errMsg ErrorMessage
	_ = json.Unmarshal(origin, &errMsg)

	return errMsg.Message
}

func (ali Alidns) InitClient() (client *alidns.Client) {
	client, err := alidns.NewClientWithAccessKey("cn-hangzhou", ali.AKID, ali.AKEY)

	if err != nil {
		logrus.Fatal(err)
	}

	return
}

// Add alidns record
// https://help.aliyun.com/document_detail/29772.html?spm=a2c4g.11186623.2.32.56eb6379m7Ssb4
func (ali Alidns) Add(subDomain string, recordType string, recordValue string) (recordId int) {
	client := ali.InitClient()
	req := alidns.CreateAddDomainRecordRequest()
	req.Scheme = "https"

	req.DomainName = ali.Domain
	req.RR = subDomain
	req.Type = recordType
	req.Value = recordValue

	response, err := client.AddDomainRecord(req)
	if err != nil {
		logrus.Fatal(getMessage(response.GetHttpContentBytes()))
		return 0
	}

	logrus.Infof("Added: %s.%s (%s %s)", subDomain, ali.Domain, recordType, recordValue)

	recordId, _ = strconv.Atoi(response.RecordId)
	return

}
