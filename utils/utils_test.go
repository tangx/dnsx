package utils

import (
	"log"
	"os"
	"testing"
)

func TestFileExists(t *testing.T) {
	a := `/tmp/fileexists.txt`
	FileExists(a)
}

func TestUserHome(t *testing.T) {
	// 1
	log.Println(UserHome())

	// 2
	log.Println(os.Getenv("HOME"))
	// 3
	home, _ := os.UserHomeDir()
	log.Println(home)
}
