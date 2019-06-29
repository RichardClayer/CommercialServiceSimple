package auth

import (
    "encoding/json"
    "fmt"
    "github.com/BiLuoHui/CommercialServiceSimple/tool/redis"
    "log"
    "strconv"
)

// 保存在线用户数据的 Redis key
const redisKey = "onlineUsers"

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
func Online(t string, u OnLineUser) (err error) {
    cfg, err := getRedisConfig()
    if err != nil {
        return
    }

    c, err := redis.New(cfg)
    if err != nil {
        return
    }
    v, err := json.Marshal(u)
    if err != nil {
        return
    }
    res, err := c.HSet(redisKey, t, string(v)).Result()
    if err != nil {
        return
    }
    if !res {
        return fmt.Errorf("保存用户登录数据失败")
    }

    return nil
}

// Offline 下线操作
func Offline(t string) (err error) {
    cfg, err := getRedisConfig()
    if err != nil {
        return
    }
    c, err := redis.New(cfg)
    if err != nil {
        return
    }
    _, err = c.HDel(redisKey, t).Result()
    if err != nil {
        return
    }

    return nil
}

// 是否已上线
func IsOnline(t string) bool {
    cfg, err := getRedisConfig()
    if err != nil {
        log.Printf("检查"+t+"是否上线，获取redis配置失败：%v\n", err)
        return false
    }
    c, err := redis.New(cfg)
    if err != nil {
        log.Printf("检查"+t+"是否上线，创建redis客户端失败：%v\n", err)
        return false
    }

    b, err := c.HExists(redisKey, t).Result()
    if err != nil {
        log.Printf("检查"+t+"是否上线，获取用户信息缓存失败：%v\n", err)
        return false
    }

    return b
}

// 线上用户信息
func UserInfo(t string) (u OnLineUser, ok bool) {
    u, ok = onlineUsers[t]
    cfg, err := getRedisConfig()
    if err != nil {
        log.Printf("获取线上用户信息，获取redis配置失败：%v\n", err)
        return OnLineUser{}, false
    }
    c, err := redis.New(cfg)
    if err != nil {
        log.Printf("获取线上用户信息，创建redis客户端失败：%v\n", err)
        return OnLineUser{}, false
    }
    s, err := c.HGet(redisKey, t).Result()
    if err != nil {
        log.Printf("获取线上用户信息失败：%v\n", err)
        return OnLineUser{}, false
    }
    if err = json.Unmarshal([]byte(s), &u); err != nil {
        log.Printf("获取线上用户信息，解析缓存数据失败：%v\n", err)
        return OnLineUser{}, false
    }
    return
}

// getRedisConfig 获取存储在线用户的 redis 配置
func getRedisConfig() (cfg redis.Config, err error) {
    dbConfig, err := redis.Get()
    if err != nil {
        return
    }
    cfg.DB, err = strconv.Atoi(dbConfig["onlineUserDB"])
    if err != nil {
        return
    }
    cfg.Addr = dbConfig["host"]
    cfg.Port = dbConfig["port"]
    cfg.Password = dbConfig["password"]

    return
}
