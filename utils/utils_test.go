package utils

import "testing"

func TestFileExists(t *testing.T) {
	a := `/tmp/fileexists.txt`
	FileExists(a)
}
