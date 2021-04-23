package record

import (
	"github.com/gin-gonic/gin"
	"github.com/tangx/dnsx/pkg/httpserver/config"
	"gorm.io/gorm"
)

var db = config.DbConn

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

func init() {
	db.AutoMigrate(&RecordItem{})
}

func RegisterRouters(rg *gin.RouterGroup) {
	record := rg.Group("/record/:record")

	record.GET("")
}
