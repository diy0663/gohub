package requests

import (
	"github.com/diy0663/go_project_packages/verifycode"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// 处理表单验证 的报错处理结果

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty"  valid:"phone"`
}
type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty"  valid:"email"`
}

func ValidateSignupPhoneExist(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"phone": []string{
			"required",
			"digits:11",
		},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
	}
	return validate(data, rules, messages)
}

func ValidateSignupEmailExist(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"email": []string{
			"required",
			"min:4",
			"max:30",
			"email",
		},
	}

	messages := govalidator.MapData{
		"email": []string{
			"required:email为必填项，参数名称 email",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
	}
	return validate(data, rules, messages)
}

type SignupUsingEmailRequest struct {
	Email           string `json:"email,omitempty"  valid:"email"`
	VerifyCode      string `json:"verify_code,omitempty"  valid:"verify_code"`
	Name            string `json:"name,omitempty"  valid:"name"`
	Password        string `json:"password,omitempty"  valid:"password"`
	PasswordConfirm string `json:"password_confirm,omitempty"  valid:"password_confirm"`
}

func SignupUsingEmail(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"email":            []string{"required", "min:4", "max:30", "email", "not_exists:users,email"},
		"verify_code":      []string{"required", "min:6"},
		"name":             []string{"required", "alpha_num", "between:3,20", "not_exists:users,name"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
	}
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
			"not_exists:Email 已被占用",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
		"name": []string{
			"required:用户名为必填项",
			"alpha_num:用户名格式错误，只允许数字和英文",
			"between:用户名长度需在 3~20 之间",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"password_confirm": []string{
			"required:确认密码框为必填项",
		},
	}
	errs := validate(data, rules, messages)
	//额外验证

	_data := data.(*SignupUsingEmailRequest)
	if _data.Password != _data.PasswordConfirm {
		errs["password"] = append(errs["password"], "两次密码输入不一致")
	}

	// 两个密码是否一致

	// 验证码是否与redis里面的一致
	if ok := verifycode.NewVerifyCode().CheckAnswer(_data.Email, _data.VerifyCode); !ok {
		errs["verify_code"] = append(errs["verify_code"], "邮件验证码不正确")
	}

	return errs
}
