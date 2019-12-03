package utils

import (
	"fmt"
	"testing"
)

func TestFileExists(t *testing.T) {
	a := `/tmp/fileexists.txt`
	FileExists(a)
}

func TestUserHome(t *testing.T) {
	fmt.Println(UserHome())
}
