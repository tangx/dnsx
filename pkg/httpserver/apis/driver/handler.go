package driver

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tangx/dnsx/pkg/httpserver/response"
)

func Create(c *gin.Context) {
	driver := Driver{}
	err := c.BindJSON(&driver)
	if err != nil {
		logrus.Errorf("bind driver failed: %v", err)
	}

	result := db.Create(&driver)
	if result.RowsAffected == 1 {
		msg := "Insert driver record success"
		response.RespOK(c, msg)
		return
	}

	if err := result.Error; err != nil {
		response.RespInternalServerError(c, 500, err)
		return
	}
}

func GetDriverByName(c *gin.Context) {

	name := c.Param("driver")
	driver, err := getDriverByName(name)
	if err != nil {
		errMsg := fmt.Errorf("get record failed: %v", err)
		response.RespInternalServerError(c, 500, errMsg)
		return
	}

	response.RespOK(c, driver)

}

func DeleteDriverByName(c *gin.Context) {
	name := c.Param("driver")
	err := deleteDriverByName(name)

	if err != nil {
		errMsg := fmt.Errorf("delete record failed: %v", err)
		response.RespInternalServerError(c, 500, errMsg)
		return
	}

	response.RespOK(c, "delete recode success")
}
