package response

const Success = 0000

// 请求类错误
const (
    AcceptNone         = iota + 1000 // 没有配置 Accept 头
    VersionNone                      // 没有检测到版本
    VersionNotSupport                // 不受支持的版本
    BadRequest                       // 无法解析的请求
    ParamsVerifyFailed               // 请求参数验证失败
)

// 认证类错误
const (
    NeedLogin = iota + 2000 // 需登录
)

// 系统类错误
const (
    DBConnectFailed = iota + 3000 // 数据库链接失败
    DBQueryFailed                 // 数据查询错误
)

// 商户相关业务错误
const (
    MerchantHasRegistered  = iota + 4000 // 商户已注册
    MerchantRegisterFailed               // 商户注册失败
)

// 员工相关业务错误
const (
    EmployeePwdErr     = iota + 5000 // 员工登录密码加密失败
    EmployeeSaveFailed               // 员工账号数据保存失败
)
