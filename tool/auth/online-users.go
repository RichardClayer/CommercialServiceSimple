package auth

// 登录认证
type OnLineUser struct {
    Id           int    `json:"id"`
    Name         string `json:"name"`
    UserName     string `json:"user_name"`
    IsRefundable int    `json:"is_refundable"`
    IsForbidden  int    `json:"is_forbidden"`
    Position     int    `json:"position"`
    Token        string `json:"token"`
}

var onlineUsers = make(map[string]OnLineUser)

// Online 上线操作 或 更新 信息
func Online(t string, u OnLineUser) {
    onlineUsers[t] = u
}

// Offline 下线操作
func Offline(t string) {
    delete(onlineUsers, t)
}

// 是否已上线
func IsOnline(t string) bool {
    _, ok := onlineUsers[t]

    return ok
}

// 线上用户信息
func UserInfo(t string) (u OnLineUser, ok bool) {
    u, ok = onlineUsers[t]
    return
}
