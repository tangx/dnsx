package driver

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func getDriverByName(name string) (*Driver, error) {
	driver := &Driver{}
	ret := db.First(driver, "name = ?", name)

	if errors.Is(ret.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("record not found")
	}

	if ret.Error != nil {
		return nil, ret.Error
	}

	return driver, nil
}
func deleteDriverByName(name string) error {
	driver, err := getDriverByName(name)
	if err != nil {
		return err
	}

	ret := db.Unscoped().Delete(driver)
	if ret.Error != nil {
		return ret.Error
	}

	return nil
}
