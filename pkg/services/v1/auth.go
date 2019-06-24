package v1

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/BiLuoHui/CommercialServiceSimple/tool/auth"
    "github.com/BiLuoHui/CommercialServiceSimple/tool/database"
    "github.com/BiLuoHui/CommercialServiceSimple/tool/request"
    "github.com/BiLuoHui/CommercialServiceSimple/tool/response"
    "golang.org/x/crypto/bcrypt"
)

// IsRegistered 是否已注册商户
func IsRegistered(w http.ResponseWriter, r *http.Request) {
    isRegistered, err := isRegistered()
    if err != nil {
        response.SendError(w, r, response.DBConnectFailed, err.Error())
        return
    }

    response.SendSuccess(w, r, map[string]bool{
        "is_registered": isRegistered,
    })
}

// Register 商户注册 仅在商户表为空时方可注册成功
func Register(w http.ResponseWriter, r *http.Request) {
    // 判断是否已注册商户
    isRegistered, err := isRegistered()
    if err != nil {
        response.SendError(w, r, response.DBConnectFailed, err.Error())
        return
    }
    if isRegistered {
        response.SendError(w, r, response.MerchantHasRegistered, "已注册商户请登录")
        return
    }
    defer r.Body.Close()

    // 获取 商户名称、登录账号、密码
    rd, err := request.GetParams(r)
    if err != nil {
        log.Printf("无法解析传参：%s\n", err)
        response.SendError(w, r, response.BadRequest, "无法解析传参")
        return
    }

    // 参数验证
    if err = registerParamsVerify(rd); err != nil {
        response.SendError(w, r, response.ParamsVerifyFailed, err.Error())
        return
    }

    // 保存商户数据
    m := Merchant{
        Name:                rd["name"],
        IsWechatPayRecorded: 0,
        IsAliPayRecorded:    0,
    }
    if err := m.Create(); err != nil {
        response.SendError(w, r, response.MerchantRegisterFailed, err.Error())
        return
    }

    // 保存登录账号数据
    pwd, err := bcrypt.GenerateFromPassword([]byte(rd["password"]), bcrypt.DefaultCost)
    if err != nil {
        log.Printf("商户登录账号密码加密失败：%v", err)
        response.SendError(w, r, response.EmployeePwdErr, "登录账号密码加密失败")
        return
    }
    e := Employee{
        Name:         rd["name"],
        UserName:     rd["username"],
        Password:     string(pwd),
        IsRefundable: 1,
        IsForbidden:  0,
        Position:     3,
        CreatedAt:    time.Now(),
        UpdatedAt:    time.Now(),
    }
    if err := e.Create(); err != nil {
        log.Printf("商户登录信息保存失败：%v", err)
        response.SendError(w, r, response.EmployeeSaveFailed, "商户登录信息保存失败")
        return
    }

    response.SendSuccess(w, r, nil)
}

// 登录信息
func Login(w http.ResponseWriter, r *http.Request) {
    // 获取 username、password
    rd, err := request.GetParams(r)
    if err != nil {
        log.Printf("无法解析传参：%s\n", err)
        response.SendError(w, r, response.BadRequest, "无法解析传参")
        return
    }

    // 参数验证
    if err = loginParamsVerify(rd); err != nil {
        response.SendError(w, r, response.ParamsVerifyFailed, err.Error())
        return
    }

    // 账号信息
    e := Employee{UserName: rd["username"]}
    if err = e.Get(); err != nil {
        response.SendError(w, r, response.LoginAccountFailed, "账号或密码不正确")
        return
    }

    // 账号是否已禁用
    if e.IsForbidden == EmployeeHasForbidden {
        response.SendError(w, r, response.AccountForbidden, "账号被禁用")
        return
    }

    // 密码验证
    err = bcrypt.CompareHashAndPassword([]byte(e.Password), []byte(rd["password"]))
    if err != nil {
        response.SendError(w, r, response.LoginAccountFailed, "账号或密码不正确")
        return
    }

    // token生成
    token, err := auth.GenToken()
    if err != nil {
        response.SendError(w, r, response.LoginFailed, "登录信息生成失败")
        return
    }

    // 保存登录信息
    onlineUser := auth.OnLineUser{
        Id:           e.Id,
        Name:         e.Name,
        UserName:     e.UserName,
        IsRefundable: e.IsRefundable,
        IsForbidden:  e.IsForbidden,
        Position:     e.Position,
        Token:        token,
    }
    auth.Online(token, onlineUser)

    response.SendSuccess(w, r, onlineUser)
}

// isRegistered 判断是否已注册商户
func isRegistered() (r bool, err error) {
    var count int
    var isRegistered bool

    c, err := database.Connect()
    if err != nil {
        return false, fmt.Errorf("数据库连接失败：%v", err)
    }
    defer c.Close()

    row := c.QueryRowContext(context.Background(), "select count(*) from merchant")
    err = row.Scan(&count)
    if err != nil {
        return false, err
    }
    if count == 0 {
        isRegistered = false
    } else {
        isRegistered = true
    }

    return isRegistered, nil
}

// registerParamsVerify 商户注册参数验证
func registerParamsVerify(d request.Params) error {
    name, ok := d["name"]
    if !ok || len(name) < 3 {
        return fmt.Errorf("商户名称长度必须大于2位")
    }

    userName, ok := d["username"]
    if !ok || len(userName) < 4 {
        return fmt.Errorf("登录账号长度必须大于3位")
    }

    password, ok := d["password"]
    if !ok || len(password) < 6 {
        return fmt.Errorf("密码长度不能小于6位")
    }

    return nil
}

// loginParamsVerify 登录参数验证
func loginParamsVerify(d request.Params) error {
    userName, ok := d["username"]
    if !ok || len(userName) == 0 {
        return fmt.Errorf("登录账号必传")
    }

    password, ok := d["password"]
    if !ok || len(password) == 0 {
        return fmt.Errorf("密码必传")
    }

    return nil
}
