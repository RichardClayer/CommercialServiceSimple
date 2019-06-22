package v1

// 商户
type Merchant struct {
	Id                  int    `json:"id"`
	Name                string `json:"name"`
	ShortName           string `json:"short_name"`
	IsWechatPayRecorded int    `json:"is_wechat_pay_recorded"`
	IsAliPayRecorded    int    `json:"is_ali_pay_recorded"`
}
