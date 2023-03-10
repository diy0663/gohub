package verifycode

// 短信数据验证码的存储,可能存在redis, 也可能存到memcached ,所以要定义一个接口, 具体用哪种存储就去实现这个接口
type Store interface {
	// 保存验证码
	Set(id string, value string) bool

	// 获取验证码
	Get(id string, clear bool) string

	// 检查验证码
	Verify(id, answer string, clear bool) bool
}
