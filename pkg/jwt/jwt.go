package jwt

import (
	"errors"
	"strings"
	"time"

	"github.com/diy0663/gohub/pkg/app"
	"github.com/diy0663/gohub/pkg/config"
	"github.com/diy0663/gohub/pkg/logger"
	"github.com/gin-gonic/gin"
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
		MaxRefresh: time.Duration(config.GetInt64("jwt.max_refresh_minute")) * time.Minute,
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

//  根据旧token 刷新得到新token ,约定了token 都是放到header里面去获取

func (jwt *JWT) RefreshToken(c *gin.Context) (string, error) {
	tokenString, err := jwt.getTokenFromHeader(c)
	if err != nil {
		return "", err
	}

	// 解析这个token ,确保这个token不是伪造的,解析出来的也可能是过期的
	token, err := jwt.parseTokenString(tokenString)
	if err != nil {
		// 假如这个报错, 不是过期的那种错误, 就不再往下走了
		validationErr, ok := err.(*jwtpkg.ValidationError)
		if !ok || validationErr.Errors != jwtpkg.ValidationErrorExpired {
			return "", err
		}
	}
	// 解析出 claims 数据, 从里面获取过期时间等来判断是否可以继续刷新延期
	claim := token.Claims.(*JWTCustomClaims)
	// 当时的签发生效日期+最大的允许刷新时间 < 当前时间, 就说明已经错过允许刷新的时间点了, 不能再帮忙延长了
	if claim.IssuedAt < app.TimenowInTimezone().Add(-jwt.MaxRefresh).Unix() {
		return "", ErrTokenExpiredMaxRefresh
	}
	claim.StandardClaims.ExpiresAt = jwt.expireAtTime()
	claim.ExpireAtTime = jwt.expireAtTime()
	return jwt.createToken(*claim)

}

//	约定token 放header
//
// 格式: Authorization:Bearer xxxxx
func (jwt *JWT) getTokenFromHeader(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrHeaderEmpty
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) == 2 && parts[0] == "Bearer" {
		return parts[1], nil
	} else {
		return "", ErrHeaderMalformed
	}
}

func (jwt *JWT) parseTokenString(tokenString string) (*jwtpkg.Token, error) {
	// 第三个匿名函数里面的 return jwt.SignKey, nil 是咋回事 ?
	return jwtpkg.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *jwtpkg.Token) (interface{}, error) {
		return jwt.SignKey, nil
	})
}

// 中间件使用 从 gin.Context 里面解析token
func (jwt *JWT) ParserToken(c *gin.Context) (*JWTCustomClaims, error) {
	tokenString, err := jwt.getTokenFromHeader(c)
	if err != nil {
		return nil, err
	}

	// 解析token串
	token, err := jwt.parseTokenString(tokenString)
	if err != nil {
		// 细化各种报错
		validationErr, ok := err.(*jwtpkg.ValidationError)
		if ok {
			if validationErr.Errors == jwtpkg.ValidationErrorMalformed {
				return nil, ErrTokenMalformed
			} else if validationErr.Errors == jwtpkg.ValidationErrorExpired {
				return nil, ErrTokenExpired
			}
		}
		return nil, ErrTokenInvalid
	}
	// token 中的 claims 信息解析出来和 JWTCustomClaims 数据结构进行校验
	if claims, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid

}
