package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/tangx/dnsx/pkg/httpserver/apis/domain"
)

var baseRouter *gin.RouterGroup

func RegisterBaseRouter(e *gin.Engine) {
	baseRouter = e.Group("dnsx")
	v0Router := baseRouter.Group("v0")

	// sub
	domain.RegisterRouters(v0Router)

}
