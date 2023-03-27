package validators

import (
	"github.com/diy0663/gohub/pkg/captcha"
	"github.com/diy0663/gohub/pkg/verifycode"
)

// 检测两个密码是否一致, 把之前存在的报错数据也传进来, 假如密码不一致,就继续把新的报错追加进去之后并再返回
func ValidatePasswordConfirm(password, passwordConfirm string, errs map[string][]string) map[string][]string {
	if password != passwordConfirm {
		errs["password_confirm"] = append(errs["password_confirm"], "两次输入密码不匹配！")
	}
	return errs
}

// ValidateVerifyCode 自定义规则，验证『手机/邮箱验证码』
func ValidateVerifyCode(key, answer string, errs map[string][]string) map[string][]string {
	if ok := verifycode.NewVerifyCode().CheckAnswer(key, answer); !ok {
		errs["verify_code"] = append(errs["verify_code"], "验证码错误")
	}
	return errs
}

func ValidateCaptcha(captchaId string, captchaAnswer string, errs map[string][]string) map[string][]string {
	if ok := captcha.NewCaptcha().Verify(captchaId, captchaAnswer, false); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "验证码错误")
	}
	return errs
}
