package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type PaginationRequest struct {
	// todo 注意, 分页参数这里是从 form 里面去验证
	Sort    string `json:"sort,omitempty" valid:"sort" form:"sort"`
	Order   string `json:"order,omitempty"  valid:"order" form:"order"`
	PerPage string `json:"per_page,omitempty" valid:"per_page" form:"per_page"`
}

func Pagination(data interface{}, c *gin.Context) map[string][]string {

	//注意,没有说必填,而是存在才进行验证
	rules := govalidator.MapData{
		// 排序字段范围
		"sort": []string{"in:id,created_at,updated_at"},
		// 升降序的范围
		"order":    []string{"in:asc,desc"},
		"per_page": []string{"numeric_between:2,100"},
	}
	messages := govalidator.MapData{
		"sort": []string{
			"in:排序字段仅支持 id,created_at,updated_at",
		},

		"order": []string{
			"in:排序规则仅支持 asc（正序）,desc（倒序）",
		},
		"per_page": []string{
			"numeric_between:每页条数的值介于 2~100 之间",
		},
	}
	return validate(data, rules, messages)
}
