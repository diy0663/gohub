package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"net/http"

	"github.com/diy0663/go_project_packages/captcha"
	"github.com/diy0663/go_project_packages/logger"
	"github.com/gin-gonic/gin"
)

type VerifyCodeController struct {
	v1.BaseAPIController
}

// 生产图形验证码的接口
func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	logger.LogIf(err)
	c.JSON(http.StatusOK, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}
