package v1

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/BiLuoHui/CommercialServiceSimple/tool/database"
    "github.com/BiLuoHui/CommercialServiceSimple/tool/response"
    "golang.org/x/crypto/bcrypt"
    "io/ioutil"
    "log"
    "net/http"
    "time"
)

type registerData map[string]string

// IsRegistered 是否已注册商户
func IsRegistered(w http.ResponseWriter, _ *http.Request) {
    isRegistered, err := isRegistered()
    if err != nil {
        response.SendError(w, response.DBConnectFailed, err.Error())
    }

    response.SendSuccess(w, map[string]bool{
        "is_registered": isRegistered,
    })
}

// Register 商户注册 仅在商户表为空时方可注册成功
func Register(w http.ResponseWriter, r *http.Request) {
    // 判断是否已注册商户
    isRegistered, err := isRegistered()
    if err != nil {
        response.SendError(w, response.DBConnectFailed, err.Error())
    }
    if isRegistered {
        response.SendError(w, response.MerchantHasRegistered, "已注册商户请登录")
    }

    // 获取 商户名称、登录账号、密码
    b, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Println("请求解析错误：" + err.Error())
        response.SendError(w, response.BadRequest, "无法解析的请求")
    }
    defer r.Body.Close()

    rd := make(registerData)
    err = json.Unmarshal(b, &rd)
    if err != nil {
        log.Printf("无法解析传参：%s\n", err)
        response.SendError(w, response.BadRequest, "无法解析传参")
    }

    // 参数验证
    if err = registerParamsVerify(rd); err != nil {
        response.SendError(w, response.ParamsVerifyFailed, err.Error())
    }

    // 保存商户数据
    m := Merchant{
        Name:                rd["name"],
        IsWechatPayRecorded: 0,
        IsAliPayRecorded:    0,
    }
    if err := m.Create(); err != nil {
        response.SendError(w, response.MerchantRegisterFailed, err.Error())
    }

    // 保存登录账号数据
    pwd, err := bcrypt.GenerateFromPassword([]byte(rd["password"]), bcrypt.DefaultCost)
    if err != nil {
        log.Printf("商户登录账号密码加密失败：%v", err)
        response.SendError(w, response.EmployeePwdErr, "登录账号密码加密失败")
    }
    e := Employee{
        Name:         rd["name"],
        UserName:     rd["username"],
        Password:     string(pwd),
        IsRefundable: 1,
        IsForbidden:  0,
        CreatedAt:    time.Now(),
        UpdatedAt:    time.Now(),
    }
    if err := e.Create(); err != nil {
        log.Printf("商户登录信息保存失败：%v", err)
        response.SendError(w, response.EmployeeSaveFailed, "商户登录信息保存失败")
    }
}

func Login(w http.ResponseWriter, r *http.Request) {
    // 获取username、password
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
    row.Scan(&count)
    if count == 0 {
        isRegistered = false
    } else {
        isRegistered = true
    }

    return isRegistered, nil
}

// registerParamsVerify 商户注册参数验证
func registerParamsVerify(d registerData) error {
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
