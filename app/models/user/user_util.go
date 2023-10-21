package user

import (
	"github.com/diy0663/gohub/pkg/app"
	"github.com/diy0663/gohub/pkg/database"
	"github.com/diy0663/gohub/pkg/paginator"
	"github.com/gin-gonic/gin"
)

// 实用工具类, 特定查询

// 判断email 是否已经被注册过了
func IsEmailExists(email string) bool {
	var count int64
	database.DB.Model(&User{}).Where("email = ? ", email).Count(&count)
	return count > 0
}

func IsPhoneExists(phone string) bool {
	var count int64
	database.DB.Model(&User{}).Where("phone = ? ", phone).Count(&count)
	return count > 0
}

func GetByPhone(phone string) (userModel User) {
	database.DB.Where("phone=?", phone).First(&userModel)
	return

}

func Get(idStr string) (userModel User) {
	database.DB.Where("id=?", idStr).First(&userModel)
	return
}

func GetByMulti(loginId string) (userModel User) {
	database.DB.Where("phone=?", loginId).Or("email=?", loginId).Or("name=?", loginId).First(&userModel)
	return
}

// 不带分页的获取全部数据..
func All() (users []User) {
	database.DB.Find(&users)
	return users
}

// 分页列表查询
func Paginate(c *gin.Context, perPage int) (users []User, paging paginator.Paging) {

	query := database.DB.Model(&User{})

	// 条件查询
	if email, isExists := c.GetQuery("email"); isExists {
		query = query.Where("email=?", string(email))
	}

	if name, isExists := c.GetQuery("name"); isExists {
		query = query.Where("name=?", string(name))
	}

	paging = paginator.Paginate(
		c,
		query,
		&users,
		app.V1URL(database.TableNameByStruct(User{})),
		perPage,
	)

	return
}
