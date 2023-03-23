package hash

import (
	"github.com/diy0663/gohub/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

// Bcrypt是单向Hash加密算法 ,不可反向破解生成明文

// 对指定字符串做哈希加密, 即便是相同的原始数据串,每次生成的也会是不一样的,所以才需要用下面的CompareHashAndPassword去验证
func BcryptHash(password string) string {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	logger.LogIf(err)
	return string(bytes)
}

// 密码对比, 第一个是明文密码, 第二个是经过加密之后存到数据库的密码
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func BcryptIsHashed(str string) bool {
	return len(str) == 60
}
