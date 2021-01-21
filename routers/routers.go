package routers
import (
	"rbac/conf"

	"github.com/gin-gonic/gin"
)

type Router struct {
	config *conf.Config
}

func NewRouters() *Router{
	result := Router{

	}

	return &result
}

func (r *Router) InitGin() *gin.Engine {
	engine := gin.Default()
	engine.GET("/health", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message": "OK",
		})
	})

	return engine
}