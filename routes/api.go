package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 在这里做 api 服务的 每一个具体的路由定义
func RegisterAPIRoutes(r *gin.Engine) {
	// 使用路由组
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"hello": "V1",
			})
		})
	}
}
