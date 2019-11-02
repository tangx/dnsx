package config

import (
	"fmt"
	"testing"
)

func Test_LoadCfg(t *testing.T) {
	config := LoadConfig()

	fmt.Println(config.Domains)
}

func Test_os(t *testing.T) {

	cfgfile := Configfile()
	fmt.Println(cfgfile)

}
