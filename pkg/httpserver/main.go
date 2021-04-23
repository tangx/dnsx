package httpserver

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tangx/dnsx/pkg/httpserver/apis"
)

var (
	server *gin.Engine

	Flags = ServerFlag{
		Addr: "127.0.0.1",
		Port: "12345",
	}
)

type ServerFlag struct {
	Addr string `flag:"addr" usage:"listen address"`
	Port string `flag:"port" usage:"listen port"`
}

func init() {
	server = gin.Default()
	apis.RegisterBaseRouter(server)
}

func Run() error {
	addr := fmt.Sprintf("%s:%s", Flags.Addr, Flags.Port)

	return server.Run(addr)

}
