package auth

import (
	"net/http"

	v1 "github.com/diy0663/gohub/app/http/controllers/api/v1"
	"github.com/diy0663/gohub/app/requests"
	"github.com/diy0663/gohub/pkg/captcha"
	"github.com/diy0663/gohub/pkg/logger"
	"github.com/diy0663/gohub/pkg/response"
	"github.com/diy0663/gohub/pkg/verifycode"
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

// 发送短信验证码, 前提是 要把 手机号, 以及上面 ShowCaptcha 的图片验证码的 captcha_id, 以及图片内的真实数字 一起发来做一次验证,不然发短信接口会被轰炸滥用
func (vc VerifyCodeController) SendUsingPhone(c *gin.Context) {
	// 验证参数 (表单验证), 验证的时候也处理 验证 captcha_id 跟传来的图片数字是否经过redis 验证
	request := requests.VerifyCodePhoneRequest{}
	if ok := requests.Validate(c, &request, requests.VerifyCodePhone); !ok {
		return
	}
	//  触发短信
	if ok := verifycode.NewVerifyCode().SendSMS(request.Phone); !ok {
		response.Abort500(c, "发送短信失败")
	}
	response.Success(c)
}

// 用于检查验证码是否正确的方法
// captcha.NewCaptcha().VerifyCaptcha()
