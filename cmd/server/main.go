package main

import (
    "fmt"
    "os"

    "github.com/BiLuoHui/CommercialServiceSimple/pkg/cmd"
)

func main() {
    // 启动服务器
    if err := cmd.RunServer(); err != nil {
        _, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
        os.Exit(1)
    }
}
