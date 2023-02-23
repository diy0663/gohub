package middlewares

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"

	"github.com/diy0663/gohub/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		// defer 延迟执行
		defer func() {
			//recover 可以中止 panic 造成的程序崩溃。它是一个只能在 defer 中发挥作用的函数，在其他作用域中调用不会发挥作用
			if err := recover(); err != nil {
				httpRequest, _ := httputil.DumpRequest(c.Request, true)
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						errStr := strings.ToLower(se.Error())
						if strings.Contains(errStr, "broken pipe") || strings.Contains(errStr, "connection reset by peer") {
							//链接中断 客户端中断连接为正常行为，不需要记录堆栈信息
							brokenPipe = true
						}
					}
				}

				if brokenPipe {

					logger.Error(c.Request.URL.Path, zap.Time("time", time.Now()), zap.Any("error", err), zap.String("request", string(httpRequest)))
					c.Error(err.(error))
					c.Abort()
					// 提前结束
					return
				}
				// 如果不是链接中断，就多记录一下堆栈信息
				logger.Error("recovery from panic", zap.Time("time", time.Now()),
					zap.Any("error", err),
					zap.String("request", string(httpRequest)),
					zap.Stack("stacktrace"),
				)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": "服务器内部错误，请稍后再试",
				})

			}
		}()
		c.Next()
	}
}
