package dnsprovider

import (
	"github.com/gin-gonic/gin"
	"github.com/tangx/dnsx/pkg/httpserver/config"
	"gorm.io/gorm"
)

type DnsProvider struct {
	gorm.Model

	Name      string
	AccessKey string // AccessKey or Token
	SecretKey string
	Comment   string

	DriverID string
}

var db = config.DbConn

func init() {
	_ = db.AutoMigrate(&DnsProvider{})
}

func RegisterRouters(rg *gin.RouterGroup) {
	dp := rg.Group("/provider/:provider")
	dp.GET("")
}
