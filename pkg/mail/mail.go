package mail

import (
	"sync"

	"github.com/diy0663/gohub/pkg/config"
)

// 本地docker 安装 mailhog
// docker run --name mailhog -p 1025:1025 -p 8025:8025 -d mailhog/mailhog
// web 页面访问: 127.0.0.1:8025

type From struct {
	Address string
	Name    string
}

// 这些个字段 都是参考 使用的第三方包
type Email struct {
	From From
	// 接收可以是多个邮箱
	To      []string
	Bcc     []string // ?
	Cc      []string // ? 抄送者?
	Subject string
	Text    []byte // 需要大量字符串处理的时候用[]byte，性能好很多,切片的结构和字符串类似，只是解除了只读限制
	HTML    []byte
}

type Mailer struct {
	Driver Driver
}

var once sync.Once
var internalMail *Mailer

func NewMail() *Mailer {
	once.Do(func() {
		internalMail = &Mailer{
			// 这里指定了要用 SMTP 驱动
			Driver: &SMTP{},
		}
	})
	return internalMail
}

// 其实新增这个方法 , 在于外部调用的时候可以少传第二个参数(邮件smtp配置),
func (mailer *Mailer) Send(email Email) bool {
	return mailer.Driver.Send(email, config.GetStringMapString("mail.smtp"))
}
