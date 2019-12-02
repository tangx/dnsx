package cmd

import (
	"fmt"
	"testing"
)

func Test_LoadConfig(t *testing.T) {
	c := Client{
		Config: "/Users/tangxin/.dnsx/config.json",
	}

	cfg := c.LoadConfig("default")
	fmt.Println(cfg.Provider)
}
