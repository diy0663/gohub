package validators

import (
	"errors"
	"fmt"
	"gohub/pkg/database"
	"strings"

	"github.com/thedevsaddam/govalidator"
)

// 在这自定义验证规则
func init() {

	govalidator.AddCustomRule("not_exists", func(field string, rule string, message string, value interface{}) error {
		// not_exists:users,email,32
		// 先把 not_exists: 去掉
		rule_data := strings.TrimPrefix(rule, "not_exists:")
		sub_data := strings.Split(rule_data, ",")
		table_name := sub_data[0]
		table_field := sub_data[1]
		exceptID := ""
		if len(sub_data) > 2 {
			exceptID = sub_data[2]
		}
		//用户传来的值强制转为字符串
		requestValue := value.(string)

		query := database.DB.Table(table_name).Where(table_field+" = ? ", requestValue)
		if len(exceptID) > 0 {
			query = query.Where("id != ? ", exceptID)
		}
		var count int64
		query.Count(&count)
		// 说明数据表中已存在重复值,验证不通过
		if count != 0 {
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("%v 已经被占用了", requestValue)
		}
		// 验证通过
		return nil
	})
}
