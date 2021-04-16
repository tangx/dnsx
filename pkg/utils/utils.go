package utils

import (
	"fmt"
	"os"
	"os/user"

	"github.com/sirupsen/logrus"
)

// PanicError 强制退出
func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}

// FileExists 判断文件是否存在
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return true, nil
}

// MustFileExists 文件必须存在
func MustFileExists(path string) {
	if ok, err := FileExists(path); !ok {
		panic(err)
	}
}

func HomeDir() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("get user home dir failed: %w", err)

	}
	return u.HomeDir, nil
}

func MustHomeDir() string {
	homedir, err := HomeDir()
	if err != nil {
		logrus.Fatalf("get home dir failed: %w", err)
	}
	return homedir
}
