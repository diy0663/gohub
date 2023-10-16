package auth

import (
	"github.com/diy0663/go_project_packages/captcha"
	"github.com/diy0663/go_project_packages/logger"
	"github.com/diy0663/go_project_packages/response"
	v1 "github.com/diy0663/gohub/app/http/controllers/v1"
	"github.com/gin-gonic/gin"
)

// 服务端生成随机验证码 random_code；
// 生成一个 captcha_id ，可以是随机字串，将 random_code 作为 captcha_id 的值存储到 redis 中，并设置过期时间为 15 分钟（过期时间可配置）；
// 生成一张 random_code 对应的验证码图片 captcha；
// 将 captcha (base64 编码) 和 captcha_id 返回给客户端；
// 客户端将 captcha 渲染为图片，展示给用户；
// 客户端将用户输出的内容 captcha_answer 和 captcha_id 传给服务器；
// 服务端使用 captcha_id 从 redis 中读取数据，将读取出来的数据和 captcha_answer 进行匹对验证。

type VerifyCodeController struct {
	v1.BaseAPIController
}

func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	id, base64Data, err := captcha.NewCaptcha().GenerateCaptcha()
	logger.LogIf(err)

	response.JSON(c, gin.H{
		"captcha_id":    id,
		"captcha_image": base64Data,
	})
}
