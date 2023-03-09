package sms

import (
	"sync"

	"github.com/diy0663/gohub/pkg/config"
)

type Message struct {
	// 第三方的短信模板code
	Template string
	// 需要传入的格式化数据
	Data map[string]string
	// Content 用来存啥?
	Content string
}

type SMS struct {
	Driver Driver
}

var once sync.Once
var internalSMS *SMS

// 这里写死了用aliyun短信
func NewSMS() *SMS {
	once.Do(func() {
		internalSMS = &SMS{

			Driver: &Aliyun{},
		}
	})
	return internalSMS
}
func (sms *SMS) Send(phone string, message Message) bool {
	return sms.Driver.Send(phone, message, config.GetStringMapString("sms.aliyun"))
}
