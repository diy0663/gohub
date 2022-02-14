package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 把api相关的具体路由放到这里,方法要大写,这样才能被bootstrap包引用

func RegisterAPIRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")

	// 这里其实就是个分组而已
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"hello": "world",
			})
		})

	}
}
