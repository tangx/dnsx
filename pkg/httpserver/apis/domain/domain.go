package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/tangx/dnsx/pkg/httpserver/apis/domain/record"
	"github.com/tangx/dnsx/pkg/httpserver/config"
	"gorm.io/gorm"
)

var db = config.DbConn

type Domain struct {
	gorm.Model

	Domain  string
	Comment string

	ProviderID int
}

func init() {
	_ = db.AutoMigrate(&Domain{})
}

func RegisterRouters(rg *gin.RouterGroup) {
	domain := rg.Group("/domain")

	// list domain
	domain.GET("")

	subrouter := domain.Group("/:domain")

	// get domain by name
	subrouter.GET("/")

	// sub routers
	record.RegisterRouters(subrouter)

}
