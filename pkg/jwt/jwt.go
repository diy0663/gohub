package jwt

import (
	"errors"
	"time"

	"github.com/diy0663/gohub/pkg/app"
	"github.com/diy0663/gohub/pkg/config"
	"github.com/diy0663/gohub/pkg/logger"
	jwtpkg "github.com/golang-jwt/jwt"
)

// 定义一组报错的变量
var (
	ErrTokenExpired           error = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh error = errors.New("令牌已过最大刷新时间")
	ErrTokenMalformed         error = errors.New("请求令牌格式有误")
	ErrTokenInvalid           error = errors.New("请求令牌无效")
	ErrHeaderEmpty            error = errors.New("需要认证才能访问！")
	ErrHeaderMalformed        error = errors.New("请求头中 Authorization 格式有误")
)

// 用了外部的包, 一般就是把那个包里面的用到结构体嵌入到自身的结构体,自身的结构体再加一些自己需要的字段
type JWTCustomClaims struct {
	jwtpkg.StandardClaims
	// 把一些能用来简单判断用户身份且不敏感的数据放进来,一般跟用户表里面的相关字段相对应
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	// 到期时间戳
	ExpireAtTime int64 `json:"expire_time"`
}

// 外部初始化的时候,一般要传值,这里要尽量减少传参,传参值最好能从config那边读取

type JWT struct {
	// 第三方包去生成token的时候,需要一个指定密钥
	SignKey []byte
	// 最大刷新时间段 , 一般从配置包里面读取, 配置里面最好从变量名上体现 时间单位
	MaxRefresh time.Duration
}

func NewJWt() *JWT {
	return &JWT{
		SignKey:    []byte(config.GetString("app.key")),
		MaxRefresh: time.Duration(config.GetInt64("wt.max_refresh_minute")) * time.Minute,
	}
}

// 外部调用,,基于指定的用户id跟名称以及密钥去 生成token , 一般在登录经过验证之后生成并返回给客户端, 客户端后面请求的时候附带到header中传来
func (jwt *JWT) IssueToken(userId, userName string) string {
	// 算出过期时间戳
	expireAtTime := jwt.expireAtTime()
	claims := JWTCustomClaims{
		StandardClaims: jwtpkg.StandardClaims{

			ExpiresAt: expireAtTime,
			// 首次签名时间
			IssuedAt: app.TimenowInTimezone().Unix(),
			//  签名颁发者
			Issuer: config.GetString("app.name"),
			//签名生效时间
			NotBefore: app.TimenowInTimezone().Unix(),
		},
		UserID:       userId,
		UserName:     userName,
		ExpireAtTime: expireAtTime,
	}

	// 生成token
	token, err := jwt.createToken(claims)
	if err != nil {
		logger.LogIf(err)
		return ""
	}

	return token
}

// 根据当前时间以及配置中过期时间间隔,算出具体的过期时间戳
func (jwt *JWT) expireAtTime() int64 {
	//
	timenow := app.TimenowInTimezone()

	// 过期分钟数
	var expireMinuteNum int64
	if config.GetBool("app.debug") {
		expireMinuteNum = config.GetInt64("jwt.debug_expire_minute")
	} else {
		expireMinuteNum = config.GetInt64("jwt.expire_minute")
	}

	expire := time.Duration(expireMinuteNum) * time.Minute
	return timenow.Add(expire).Unix()
}

// 使用第三包传入 JWTCustomClaims 去 生成token
func (jwt *JWT) createToken(claims JWTCustomClaims) (string, error) {
	token := jwtpkg.NewWithClaims(jwtpkg.SigningMethodHS256, claims)
	return token.SignedString(jwt.SignKey)
}
