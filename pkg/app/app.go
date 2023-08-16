package app

import (
	"gin-rest-tool/pkg/errSys"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		Ctx: ctx,
	}
}

func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) Success(data interface{}) {
	response := gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"result":  data,
	}
	r.Ctx.JSON(http.StatusOK, response)
}

func (r *Response) Error(err interface{}) {
	if v, ok := err.(*errSys.Error); ok {
		response := gin.H{
			"status":  v.Code(),
			"message": v.Msg(),
			"result":  gin.H{},
		}
		r.Ctx.JSON(http.StatusOK, response)
		return
	}

	if v, ok := err.(error); ok {
		response := gin.H{
			"status":  errSys.InvalidParams.Code(),
			"message": v.Error(),
			"result":  gin.H{},
		}
		r.Ctx.JSON(http.StatusOK, response)
		return
	}
}
