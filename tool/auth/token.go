package auth

import (
    "fmt"
    "net/http"

    "github.com/google/uuid"
)

// GenToken 生成登录token
func GenToken() (u string, err error) {
    id, err := uuid.NewUUID()
    if err != nil {
        return "", err
    }

    return id.String(), nil
}

// GetToken 获取登录token
func GetToken(r *http.Request) (t string, err error) {
    t = r.Header.Get("Authorization")
    if len(t) == 0 {
        return "", fmt.Errorf("未登录或登录已过期")
    }

    return
}
