package auth

import (
	"net/http"

	v1 "github.com/diy0663/gohub/app/http/controllers/api/v1"
	"github.com/diy0663/gohub/pkg/captcha"
	"github.com/diy0663/gohub/pkg/logger"
	"github.com/gin-gonic/gin"
)

type VerifyCodeController struct {
	v1.BaseApiController
}

func (vc VerifyCodeController) ShowCaptcha(c *gin.Context) {
	id, b64s, err := captcha.NewCaptcha().Generate()
	logger.LogIf(err)
	c.JSON(http.StatusOK, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}

// 用于检查验证码是否正确的方法
// captcha.NewCaptcha().VerifyCaptcha()
