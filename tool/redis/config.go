package redis

import (
    "github.com/Unknwon/goconfig"
)

const configFile = "config/app.ini"

// 获取Redis配置
func Get() (dbConfig map[string]string, err error) {
    conf, err := goconfig.LoadConfigFile(configFile)
    if err != nil {
        return
    }

    return conf.GetSection("redis")
}
