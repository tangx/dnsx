package driver

import (
	"github.com/gin-gonic/gin"
	"github.com/tangx/dnsx/pkg/httpserver/config"
	"gorm.io/gorm"
)

type Driver struct {
	gorm.Model

	Name    string
	Comment string
}

func init() {
	_ = config.DbConn.AutoMigrate(&Driver{})
}

func RegisterRouters(rg *gin.RouterGroup) {
	driver := rg.Group("dirver")

	// list
	driver.GET("")

	// create driver by name
	driver.POST("/:dirver")

	// list by name
	driver.GET("/:driver")

	// delete by name
	driver.DELETE("/:driver")

}
