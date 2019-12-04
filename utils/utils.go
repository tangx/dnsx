package utils

import (
	"os"
	"os/user"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func IsError(err error) {
	if err != nil {
		panic(err)
	}
}

func UserHome() string {
	u, err := user.Current()
	IsError(err)
	return u.HomeDir
}
