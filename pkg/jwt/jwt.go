package jwt

import (
	"errors"
	"strings"
	"time"

	"github.com/diy0663/go_project_packages/config"
	"github.com/diy0663/go_project_packages/logger"
	"github.com/diy0663/gohub/pkg/app"
	"github.com/gin-gonic/gin"
	jwtpkg "github.com/golang-jwt/jwt"
)

var (
	ErrTokenExpired           error = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh error = errors.New("令牌已过最大刷新时间")
	ErrTokenMalformed         error = errors.New("请求令牌格式有误")
	ErrTokenInvalid           error = errors.New("请求令牌无效")
	ErrHeaderEmpty            error = errors.New("需要认证才能访问！")
	ErrHeaderMalformed        error = errors.New("请求头中 Authorization 格式有误")
)

type JWT struct {
	SignKey    []byte
	MaxRefresh time.Duration
}

type JWTCustomClaims struct {
	UserID       string `json:"user_id,omitempty" `
	UserName     string `json:"user_name,omitempty" `
	ExpireAtTime int64  `json:"expire_at_time,omitempty" `
	jwtpkg.StandardClaims
}

// 这种一般包内很少有函数,都是写方法跟结构体进行绑定
func NewJWT() *JWT {
	return &JWT{
		SignKey:    []byte(config.GetString("app.key")),
		MaxRefresh: time.Duration(config.GetInt64("jwt.max_refresh_minute")) * time.Minute,
	}
}

// 登录的时候签发token
func (jwt *JWT) IssueToken(userID string, userName string) string {

	expireTime := jwt.ExpireAtTime()
	claims := JWTCustomClaims{

		UserID:       userID,
		UserName:     userName,
		ExpireAtTime: expireTime,
		StandardClaims: jwtpkg.StandardClaims{
			NotBefore: app.TimenowInTimezone().Unix(), // 签名生效时间
			IssuedAt:  app.TimenowInTimezone().Unix(), // 首次签名时间（后续刷新 Token 不会更新）
			ExpiresAt: expireTime,                     // 签名过期时间
			Issuer:    config.GetString("app.name"),   // 签名颁发者
		},
	}
	token, err := jwt.createToken(claims)
	if err != nil {
		logger.LogIf(err)
		return ""
	}
	return token
}

// 解析token (中间件使用)
func (jwt *JWT) ParserToken(c *gin.Context) (*JWTCustomClaims, error) {
	tokenString, err := jwt.getTokenFromHeader(c)
	if err != nil {
		return nil, err
	}

	// 库解析用户传参的 Token
	tokenData, err := jwt.parseTokenString(tokenString)
	if err != nil {
		// 区分报错
		validationErr, ok := err.(*jwtpkg.ValidationError)
		if ok {
			if validationErr.Errors == jwtpkg.ValidationErrorMalformed {
				// 格式错误
				return nil, ErrTokenMalformed
			} else if validationErr.Errors == jwtpkg.ValidationErrorExpired {
				// 过期
				return nil, ErrTokenExpired
			}
		}
		// 剩下的就是无效token
		return nil, ErrTokenInvalid
	}
	// 与 JWTCustomClaims 数据结构进行校验 并且能解析为我们需要的claims
	if claims, ok := tokenData.Claims.(*JWTCustomClaims); ok && tokenData.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid

}

// 刷新token
func (jwt *JWT) RefreshToken(c *gin.Context) (string, error) {
	tokenString, err := jwt.getTokenFromHeader(c)
	if err != nil {
		return "", err
	}

	// 解析token
	token, err := jwt.parseTokenString(tokenString)
	if err != nil {
		validationErr, ok := err.(*jwtpkg.ValidationError)
		if !ok || validationErr.Errors != jwtpkg.ValidationErrorExpired {
			// 这个报错, 不是 jwt里面的解析报错 或者 这个报错不是过期的报错,就直接返回
			return "", err
		}

	}
	// token刷新机制不是JWT标准规范的一部分，它是根据具体应用的需求自行实现的
	// 过期后,还可以允许刷新token, 但是不能超过 最大允许刷新的时间 ??

	claims := token.Claims.(*JWTCustomClaims)

	x := app.TimenowInTimezone().Add(-jwt.MaxRefresh).Unix()
	// 签发时间+最大过期时间 > 当前 ,就允许刷新
	if claims.IssuedAt > x {
		claims.StandardClaims.ExpiresAt = jwt.ExpireAtTime()
		return jwt.createToken(*claims)
	}
	return "", ErrTokenExpiredMaxRefresh
}

// 调用包解析token
func (jwt *JWT) parseTokenString(tokenString string) (*jwtpkg.Token, error) {
	return jwtpkg.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *jwtpkg.Token) (interface{}, error) {
		return jwt.SignKey, nil
	})
}

func (jwt *JWT) createToken(claims JWTCustomClaims) (string, error) {
	token := jwtpkg.NewWithClaims(jwtpkg.SigningMethodHS256, claims)
	return token.SignedString(jwt.SignKey)
}

// 计算过期时间戳
func (jwt *JWT) ExpireAtTime() int64 {
	// 根据时区获取当前时间戳

	timenow := app.TimenowInTimezone()

	var expire time.Duration
	if config.GetBool("app.debug") {
		expire = time.Duration(config.GetInt64("jwt.debug_expire_minute")) * time.Minute
	} else {
		expire = time.Duration(config.GetInt64("jwt.expire_minute")) * time.Minute
	}
	return timenow.Add(expire).Unix()

}

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
