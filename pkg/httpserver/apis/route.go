package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/tangx/dnsx/pkg/httpserver/apis/dnsprovider"
	"github.com/tangx/dnsx/pkg/httpserver/apis/domain"
	"github.com/tangx/dnsx/pkg/httpserver/apis/driver"
)

var baseRouter *gin.RouterGroup

func RegisterBaseRouter(e *gin.Engine) {
	baseRouter = e.Group("dnsx")
	v0Router := baseRouter.Group("v0")

	// sub
	driver.RegisterRouters(v0Router)
	dnsprovider.RegisterRouters(v0Router)
	domain.RegisterRouters(v0Router)

}
