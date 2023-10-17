package requests

import (
	"errors"
	"fmt"
	"strings"

	requestsPkg "github.com/diy0663/go_project_packages/requests"
	"github.com/diy0663/gohub/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

func init() {

	// 自定义规则 not_exists，验证请求数据必须不存在于数据库中。
	// 常用于保证数据库某个字段的值唯一，如用户名、邮箱、手机号、或者分类的名称。
	// not_exists 参数可以有两种，一种是 2 个参数，一种是 3 个参数：
	// not_exists:users,email 检查数据库表里是否存在同一条信息
	// not_exists:users,email,32 排除用户掉 id 为 32 的用户

	govalidator.AddCustomRule("not_exists", func(field string, rule string,
		message string, value interface{}) error {

		parameter := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")
		tableName := parameter[0]
		dbFiled := parameter[1]
		var exceptID string
		if len(parameter) > 2 {
			exceptID = parameter[2]
		}

		requestValue := value.(string)

		query := database.DB.Table(tableName).Where(dbFiled+" = ? ", requestValue)
		if len(parameter) > 2 {
			query.Where(dbFiled+" != ? ", exceptID)
		}

		var count int64
		query.Count(&count)
		if count != 0 {
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("%v已经被占用了", requestValue)
		}

		return nil
	})
}

func RequestValidate(c *gin.Context, obj interface{}, handler requestsPkg.ValidatorFunc) bool {
	return requestsPkg.ValidateInAPI(c, obj, handler)
}

func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	return requestsPkg.ValidateInRequest(data, rules, messages)
}
