package main

import (
    "fmt"
    "log"
    "os"

    "github.com/BiLuoHui/CommercialServiceSimple/pkg/cmd"
)

func main() {
    // 启动服务器
    if err := cmd.RunServer(); err != nil {
        _, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
        os.Exit(1)
    }

    log.Println("HTTP服务器启动成功")
}
