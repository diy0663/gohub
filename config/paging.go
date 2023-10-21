package config

import (
	configPkg "github.com/diy0663/go_project_packages/config"
)

func init() {
	configPkg.Add("paging", map[string]interface{}{

		"perpage": 10,
		// URL 中用以分辨多少页的参数
		// 此值若修改需一并修改请求验证规则
		"url_query_page": "page",

		// URL 中用以分辨排序的参数（使用 id 或者其他）
		// 此值若修改需一并修改请求验证规则
		"url_query_sort": "sort",

		// URL 中用以分辨排序规则的参数（辨别是正序还是倒序）
		// 此值若修改需一并修改请求验证规则
		"url_query_order": "order",

		// URL 中用以分辨每页条数的参数
		// 此值若修改需一并修改请求验证规则
		"url_query_per_page": "per_page",
	})
}
