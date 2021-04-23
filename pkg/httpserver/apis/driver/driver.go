package driver

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tangx/dnsx/pkg/httpserver/config"
	"gorm.io/gorm"
)

type Driver struct {
	gorm.Model

	Name    string `gorm:"uniqueIndex"`
	Comment string
}

var db = config.DbConn

func init() {
	err := db.AutoMigrate(&Driver{})
	if err != nil {
		logrus.Fatalf("auto migrate driver table failed: %v", err)
	}
}

func RegisterRouters(rg *gin.RouterGroup) {
	driver := rg.Group("dirver")

	// list
	driver.GET("")

	driver.POST("", Create)
	driver.GET("/:driver", GetDriverByName)
	driver.DELETE("/:driver", DeleteDriverByName)

}
