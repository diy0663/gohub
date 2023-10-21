package requests

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

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
	// not_exists:users,email,id<>32 排除用户掉 id 为 32 的用户

	govalidator.AddCustomRule("not_exists", func(field string, rule string,
		message string, value interface{}) error {

		parameter := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")
		tableName := parameter[0]
		dbFiled := parameter[1]
		// var exceptStr string
		// if len(parameter) > 2 {

		// 	exceptStr = parameter[2]
		// }

		requestValue := value.(string)

		query := database.DB.Table(tableName).Where(dbFiled+" = ? ", requestValue)
		if len(parameter) > 2 {
			exceptStr := parameter[2]
			exceptData := strings.SplitN(exceptStr, "<>", 2)
			exceptField, exceptValue := exceptData[0], exceptData[1]
			// 这里写死数据库字段为
			query.Where(exceptField+" != ? ", exceptValue)
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

	govalidator.AddCustomRule("max_cn", func(field string, rule string, message string, value interface{}) error {
		valLength := utf8.RuneCountInString(value.(string))
		lengthNum, _ := strconv.Atoi(strings.TrimPrefix(rule, "max_cn:"))
		if valLength > lengthNum {
			// 如果有自定义错误消息的话，使用自定义消息
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("长度不能超过 %v 个字", lengthNum)

		}

		return nil
	})

	govalidator.AddCustomRule("min_cn", func(field string, rule string, message string, value interface{}) error {
		valLength := utf8.RuneCountInString(value.(string))
		lengthNum, _ := strconv.Atoi(strings.TrimPrefix(rule, "min_cn:"))
		if valLength < lengthNum {
			// 如果有自定义错误消息的话，使用自定义消息
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("长度至少要 %v 个字", lengthNum)

		}

		return nil
	})

	//确保数据库存在某条数据
	govalidator.AddCustomRule("exists", func(field string, rule string, message string, value interface{}) error {
		rng := strings.Split(strings.TrimPrefix(rule, "exists:"), ",")

		// 第一个参数，表名称，如 categories
		tableName := rng[0]
		// 第二个参数，字段名称，如 id
		dbFiled := rng[1]

		// 用户请求过来的数据
		requestValue := value.(string)

		// 查询数据库
		var count int64
		database.DB.Table(tableName).Where(dbFiled+" = ?", requestValue).Count(&count)
		// 验证不通过，数据不存在
		if count == 0 {
			// 如果有自定义错误消息的话，使用自定义消息
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("%v 不存在", requestValue)
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
