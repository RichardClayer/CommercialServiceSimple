package auth

import (
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
