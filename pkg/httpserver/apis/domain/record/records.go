package record

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tangx/dnsx/pkg/httpserver/config"
	"gorm.io/gorm"
)

var db = config.DbConn

type Record struct {
	gorm.Model

	// ID DnsProvider 提供的
	RecordID string `gorm:"unique"`
	Record   string
	Type     string
	Value    string
	Status   bool
	Comment  string

	DomainID string
}

func init() {
	err := db.AutoMigrate(&Record{})
	if err != nil {
		logrus.Fatalf("auto migrate record tabel failed: %v", err)
	}
}

func RegisterRouters(rg *gin.RouterGroup) {
	record := rg.Group("/record")

	subrouter := record.Group("/:record")
	subrouter.GET("")
}
