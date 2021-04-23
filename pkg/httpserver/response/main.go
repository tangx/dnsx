package response

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code  int
	Error string
	Data  interface{}
}

func RespCommon(c *gin.Context, code int, data interface{}, err error) {
	msg := ""
	if err != nil {
		msg = err.Error()
	}
	resp := Response{
		Code:  code,
		Error: msg,
		Data:  data,
	}

	c.JSON(code, resp)
}

func RespOK(c *gin.Context, data interface{}) {
	resp := Response{
		Code: 0,
		Data: data,
	}
	c.JSON(200, resp)
}

func RespInternalServerError(c *gin.Context, code int, err error) {
	msg := ""
	if err != nil {
		msg = err.Error()
	}
	resp := Response{
		Code:  code,
		Error: msg,
	}
	c.JSON(500, resp)

}
