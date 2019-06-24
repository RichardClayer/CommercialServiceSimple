package v1

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/BiLuoHui/CommercialServiceSimple/tool/database"
)

// 商户
type Merchant struct {
    Id                  int    `json:"id"`
    Name                string `json:"name"`
    IsWechatPayRecorded int    `json:"is_wechat_pay_recorded"`
    IsAliPayRecorded    int    `json:"is_ali_pay_recorded"`
}

// 保存商户
func (m *Merchant) Create() error {
    c, err := database.Connect()
    if err != nil {
        return fmt.Errorf("数据库链接失败：%v", err)
    }
    defer c.Close()
    res, err := c.ExecContext(context.Background(),
        "insert into merchant(name, is_wechat_pay_recorded, is_ali_pay_recorded) values(?, ?, ?)",
        m.Name,
        m.IsWechatPayRecorded,
        m.IsAliPayRecorded)
    if err != nil {
        return fmt.Errorf("保存商户信息失败：%v", err)
    }
    _, err = res.LastInsertId()
    if err != nil {
        return fmt.Errorf("获取保存的商户信息失败：%v", err)
    }

    return nil
}

const (
    EmployeeHasForbidden = 1 // 账号被禁用
)

// 员工
type Employee struct {
    Id           int    `json:"id"`
    Name         string `json:"name"`
    UserName     string `json:"user_name"`
    Password     string
    IsRefundable int       `json:"is_refundable"`
    IsForbidden  int       `json:"is_forbidden"`
    Position     int       `json:"position"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}

// 保存员工信息
func (e *Employee) Create() error {
    c, err := database.Connect()
    if err != nil {
        return fmt.Errorf("数据库链接失败：%v", err)
    }
    defer c.Close()
    res, err := c.ExecContext(context.Background(),
        "insert into employees(name, username, password, is_refundable, is_forbidden,position, created_at, updated_at) values(?, ?, ?, ?, ?, ?, ?, ?)",
        e.Name,
        e.UserName,
        e.Password,
        e.IsRefundable,
        e.IsForbidden,
        e.Position,
        e.CreatedAt,
        e.UpdatedAt)
    if err != nil {
        return fmt.Errorf("保存商户登录信息失败：%v", err)
    }
    _, err = res.LastInsertId()
    if err != nil {
        return fmt.Errorf("获取保存的商户登录信息失败：%v", err)
    }

    return nil
}

// Get 获取员工信息
func (e *Employee) Get() error {
    c, err := database.Connect()
    if err != nil {
        return fmt.Errorf("数据库连接失败：%v", err)
    }
    defer c.Close()

    row := c.QueryRowContext(context.Background(),
        "select id, name,username,password,is_refundable,is_refundable,position,created_at,updated_at from employees where username = ?", e.UserName)
    if err = row.Scan(&e.Id, &e.Name, &e.UserName, &e.Password, &e.IsRefundable, &e.IsRefundable, &e.Position, &e.CreatedAt, &e.UpdatedAt); err != nil {
        log.Printf("查询账号登录信息失败：%v\n", err)
        return err
    }

    return nil
}
