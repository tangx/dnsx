package utils

import "os"

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
