package validators

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/diy0663/gohub/pkg/database"
	"github.com/thedevsaddam/govalidator"
)

func init() {

	// todo  field参数没被用到
	govalidator.AddCustomRule("not_exists", func(field string, rule string, message string, value interface{}) error {

		//把 规则里面去掉not_exists: 这个前缀之后剩余的部分用, 分隔, 来获取对应的值, 在这类按照第一个值是数据表, 第二个值是字段名,第三个值是查询排除值, id != ?
		explode_fields := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")
		tableName := explode_fields[0]
		dbFields := explode_fields[1]
		var exceptId string
		if len(explode_fields) > 2 {
			exceptId = explode_fields[2]
		}
		requestValue := value.(string)
		query := database.DB.Table(tableName).Where(dbFields+" = ? ", requestValue)
		if len(exceptId) > 0 {
			query.Where("id != ? ", exceptId)
		}
		var count int64
		query.Count(&count)
		if count != 0 {
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("%v 已被占用", requestValue)
			// return errors.New(requestValue + " 已经被占用")
		}

		return nil
	})

	// 中文字符的最大长度限制
	govalidator.AddCustomRule("max_cn", func(field string, rule string, message string, value interface{}) error {
		valLength := utf8.RuneCountInString(value.(string))
		l, _ := strconv.Atoi(strings.TrimPrefix(rule, "max_cn:"))
		if valLength > l {
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("长度不能超过 %d 个字", l)
		}

		return nil
	})

	govalidator.AddCustomRule("min_cn", func(field string, rule string, message string, value interface{}) error {

		valLength := utf8.RuneCountInString(value.(string))
		l, _ := strconv.Atoi(strings.TrimPrefix(rule, "min_cn:"))
		if valLength < l {
			// 如果有自定义错误消息的话，使用自定义消息
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("长度需大于 %d 个字", l)
		}
		return nil
	})

}
