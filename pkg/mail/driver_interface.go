package mail

type Driver interface {
	// 实现发送邮件
	// 入参 : 邮件内容  , 配置
	Send(email Email, config map[string]string) bool
}
