package middlewares

import (
	"fmt"

	"github.com/diy0663/go_project_packages/config"
	"github.com/diy0663/go_project_packages/response"
	"github.com/diy0663/gohub/app/models/user"
	casbinpkg "github.com/diy0663/gohub/pkg/casbinPkg"
	"github.com/diy0663/gohub/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func CasbinRule() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := jwt.NewJWT().ParserToken(c)
		if err != nil {
			response.Unauthorized(c, fmt.Sprintf("请查看 %v 相关的接口认证文档", config.GetString("app.name")))
			return
		}
		// 解析成功
		userModel := user.Get(claims.UserID)
		if userModel.ID == 0 {
			response.Unauthorized(c, "找不到对应用户，用户可能已删除")
			return
		}

		path := c.Request.URL.Path
		act := c.Request.Method
		RoleId := userModel.GetStringRoleID()
		success, _ := casbinpkg.GetCasbinRBAC().Enforce(RoleId, path, act)
		if !success {
			response.Unauthorized(c, "角色权限不足")
			return
		}

		c.Next()
	}
}
