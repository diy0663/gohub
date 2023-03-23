package requests

import (
	"github.com/diy0663/gohub/app/requests/validators"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	// omitempty 不存在的时候就不输出
	Phone string `json:"phone,omitempty" valid:"phone"` //  valid 对应 TagIdentifier ,表明出现关键字 valid 就验证,验证规则就是 valid 后面跟着的规则
}

func ValidateSignupPhoneExist(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
	}
	return validate(data, rules, messages)

}

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
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
	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
	}
	return validate(data, rules, messages)

}

// 通过手机号进行注册的请求信息
type SignupUsingPhoneRequest struct {
	Phone           string `json:"phone,omitempty"  valid:"phone" `
	VerifyCode      string `json:"verify_code,omitempty" valid:"verify_code"`
	Name            string `valid:"name" json:"name"`
	Password        string `valid:"password" json:"password,omitempty"`
	PasswordConfirm string `valid:"password_confirm" json:"password_confirm,omitempty"`
}

func SignupUsingPhone(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		// 这里 新增了一个规则 not_exists 后面接3个参数, 一个是数据表, 一个是对应的字段,用于判断某个数据不能重复
		"phone":            []string{"required", "digits:11", "not_exists:users,phone"},
		"name":             []string{"required", "alpha_num", "between:3,20", "not_exists:users,name"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
		"verify_code":      []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
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
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
	}
	errs := validate(data, rules, messages)
	// 增加其他验证 , 验证2次密码是否一致, 验证短信验证码是否与redis里面存的一致,不加这两个的话,无法引用validators,将会导致上面 not_exists 无用使用,除非 把not_exists 放到别的已经被加载进来的包里面去
	_data := data.(*SignupUsingPhoneRequest)
	errs = validators.ValidatePasswordConfirm(_data.Password, _data.PasswordConfirm, errs)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)

	return errs

}

// 通过邮箱进行注册的请求信息
type SignupUsingEmailRequest struct {
	Email           string `json:"email,omitempty"  valid:"email" `
	VerifyCode      string `json:"verify_code,omitempty" valid:"verify_code"`
	Name            string `valid:"name" json:"name"`
	Password        string `valid:"password" json:"password,omitempty"`
	PasswordConfirm string `valid:"password_confirm" json:"password_confirm,omitempty"`
}

func SignupUsingEmail(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		// 这里 新增了一个规则 not_exists 后面接3个参数, 一个是数据表, 一个是对应的字段,用于判断某个数据不能重复
		"email":            []string{"required", "email", "min:4", "max:30", "not_exists:users,email"},
		"name":             []string{"required", "alpha_num", "between:3,20", "not_exists:users,name"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
		"verify_code":      []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
			"not_exists:Email 已被占用",
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
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
	}
	errs := validate(data, rules, messages)
	// 增加其他验证 , 验证2次密码是否一致, 验证短信验证码是否与redis里面存的一致,不加这两个的话,无法引用validators,将会导致上面 not_exists 无用使用,除非 把not_exists 放到别的已经被加载进来的包里面去
	_data := data.(*SignupUsingEmailRequest)
	errs = validators.ValidatePasswordConfirm(_data.Password, _data.PasswordConfirm, errs)
	errs = validators.ValidateVerifyCode(_data.Email, _data.VerifyCode, errs)

	return errs

}
