package auth

import (
	v1 "github.com/diy0663/gohub/app/http/controllers/api/v1"
	"github.com/diy0663/gohub/pkg/captcha"
	"github.com/diy0663/gohub/pkg/logger"
	"github.com/diy0663/gohub/pkg/response"
	"github.com/gin-gonic/gin"
)

type VerifyCodeController struct {
	v1.BaseAPIController
}

func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	id, base64_value, err := captcha.NewCaptcha().GenerateCaptcha()
	logger.LogIf(err)

	response.JSON(c, gin.H{
		"captcha_id":    id,
		"captcha_image": base64_value,
	})

}
