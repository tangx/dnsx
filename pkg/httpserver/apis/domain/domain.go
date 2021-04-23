package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/tangx/dnsx/pkg/httpserver/apis/domain/record"
)

func RegisterRouters(rg *gin.RouterGroup) {
	domain := rg.Group("/domain/:domain")

	domain.GET("")

	// sub routers
	record.RegisterRouters(domain)

}
