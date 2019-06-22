package response

const Success    = 0000

// 请求类错误
const (
	AcceptNone = iota + 1000
	VersionNone
)

// 认证类错误
const (
	NeedLogin = iota + 2000
)
