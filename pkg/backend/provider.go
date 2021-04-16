package backend

import (
	"github.com/tangx/dnsx/pkg/backend/internal/response"
)

type Provider interface {
	AddRecord(domain, record, rrType, Value string) (recordID string)
	GetRecords(domain, record string) (RRs []response.RecordItem)
	DeleteRecord(domain, recordID string) string
	SetRecordStatus(domain string, recordID string, status bool) string
}
