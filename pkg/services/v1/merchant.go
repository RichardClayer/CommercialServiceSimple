package v1

import (
    "context"
    "fmt"
    "github.com/BiLuoHui/CommercialServiceSimple/tool/database"
)

// 商户基础资料管理、微信商户配置、支付宝商户配置
type WechatPayConfig struct {
    Id            int    `json:"id"`
    AppId         string `json:"app_id"`
    AppSecret     string `json:"app_secret"`
    MchId         string `json:"mch_id"`
    Key           string `json:"key"`
    APIClientCert string `json:"api_client_cert"`
    APIClientKey  string `json:"api_client_key"`
    MerchantId    int    `json:"merchant_id"`
}

// Get 获取微信支付配置
func (w *WechatPayConfig) Get() error {
    c, err := database.Connect()
    if err != nil {
        return fmt.Errorf("数据库链接失败：%v", err)
    }
    defer c.Close()
    row := c.QueryRowContext(context.Background(), "select * from wechat_pay limit 1")
    if err = row.Scan(&w.Id, &w.AppId, &w.AppSecret, &w.MchId, &w.Key, &w.APIClientCert, &w.APIClientKey, &w.MerchantId); err != nil {
        return fmt.Errorf("获取微信支付配置失败：%v", err)
    }

    return nil
}
