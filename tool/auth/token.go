package auth

import (
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
func GetToken(r *http.Request) string {
    return r.Header.Get("Authorization")
}
