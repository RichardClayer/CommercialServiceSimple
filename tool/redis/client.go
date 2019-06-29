package redis

import (
    "github.com/go-redis/redis"
)

// Redis 配置
type Config struct {
    Addr     string
    Port     string
    Password string
    DB       int
}

// New 创建新的 Redis 客户端
func New(c Config) (client *redis.Client, err error) {
    client = redis.NewClient(&redis.Options{
        Addr:     c.Addr + ":" + c.Port,
        Password: c.Password,
        DB:       c.DB,
    })
    _, err = client.Ping().Result()
    if err != nil {
        return &redis.Client{}, err
    }

    return client, nil
}
