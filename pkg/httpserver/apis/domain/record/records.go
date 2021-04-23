package record

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RecordItem struct {
	gorm.Model

	// ID DnsProvider 提供的
	RecordID string `gorm:"unique"`
	Domain   string
	Record   string
	Type     string
	Value    string
	Status   bool
	Comment  string
}

func RegisterRouters(rg *gin.RouterGroup) {
	record := rg.Group("/record/:record")

	record.GET("")
}
