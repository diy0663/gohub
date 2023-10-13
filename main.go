package main

import (
	"fmt"

	config "github.com/diy0663/go_project_packages/config"
	"github.com/diy0663/gohub/bootstrap"
	btsConfig "github.com/diy0663/gohub/config"
	"github.com/gin-gonic/gin"
)

func init() {
	btsConfig.InitAllConfig()
}

// main  入口保持整洁
func main() {
	// cmd+shift+K 用于vscode 删除 单行
	// 自行配置 cmd+D 用于复制单行

	r := gin.New()
	// 中间件
	bootstrap.SetupRoute(r)
	//  http://127.0.0.1:8080
	err := r.Run(":" + config.GetString("app.port"))
	if err != nil {
		//  错误处理，端口被占用了或者其他错误
		fmt.Println(err.Error())
	}

}
