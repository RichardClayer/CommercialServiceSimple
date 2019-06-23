package response

const Success = 0000

// 请求类错误
const (
    AcceptNone = iota + 1000
    VersionNone
    VersionNotMatch
    VersionNotSupport
)

// 认证类错误
const (
    NeedLogin = iota + 2000
)

// 系统类错误
const (
    DBConnectFailed = iota + 3000
    DBQueryFailed
)
