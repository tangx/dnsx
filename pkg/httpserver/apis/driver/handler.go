package driver

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Create(c *gin.Context) {
	driver := Driver{}
	err := c.BindJSON(&driver)
	if err != nil {
		logrus.Errorf("bind driver failed: %v", err)
	}

	tx := db.Create(&driver)
	if tx.RowsAffected == 1 {
		// todo: gin render ok
		c.String(200, "driver record insert success")
		return
	}

	if err := tx.Error; err != nil {
		logrus.Errorf("dirver insert failed: %v", err)
		// todo: gin render 50x
		return
	}
}

func GetDriverByName(c *gin.Context) {

	name := c.Param("driver")
	driver, err := getDriverByName(name)
	if err != nil {
		// todo: gin render 50x not found
		return
	}

	c.JSON(200, driver)

}

func getDriverByName(name string) (*Driver, error) {
	driver := &Driver{}
	ret := db.First(driver, "name = ?", name)

	if ret.Error != nil {
		return nil, ret.Error
	}

	return driver, nil
}

func DeleteDriverByName(c *gin.Context) {
	name := c.Param("driver")
	err := deleteDriverByName(name)

	if err != nil {
		c.String(500, err.Error())
		return
	}

	// todo: render ok
	c.String(201, "delete ok")

}

func deleteDriverByName(name string) error {
	driver, err := getDriverByName(name)
	if err != nil {
		return nil
	}

	ret := db.Unscoped().Delete(driver)
	if ret.Error != nil {
		return ret.Error
	}

	return nil

}
