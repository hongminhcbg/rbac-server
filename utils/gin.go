package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"rbac/dtos"
)

func Response(ctx *gin.Context, code int, meta dtos.MetaData, data interface{}) {
	ctx.JSON(code, dtos.Response{
		Meta: meta,
		Data: data,
	})
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, dtos.Response{
		Meta: dtos.MetaData{
			Code: http.StatusOK,
		},
		Data: data,
	})
}

func ResponseError(ctx *gin.Context, code int, err error) {
	ctx.JSON(code, dtos.Response{
		Meta: dtos.MetaData{
			Code:    code,
			Message: err.Error(),
		},
	})
}
