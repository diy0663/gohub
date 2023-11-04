package token

import (
	"fmt"

	"github.com/diy0663/gohub/pkg/database"
)

// todo 保存之后自动import
// util 里面存的都是直接对数据表的查询操作, 是函数, 不是结构体实现的方法

func Get(idstr string) (token Token) {
	database.DB.Where("id", idstr).First(&token)
	return
}

func GetBy(field, value string) (token Token) {
	str := fmt.Sprintf("%v= ?", field)
	database.DB.Where(str, value).First(&token)
	return
}

func All() (tokens []Token) {
	database.DB.Find(&tokens)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Token{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}
