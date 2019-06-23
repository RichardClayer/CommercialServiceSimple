package cmd

import (
	"context"
	"flag"
	"fmt"

	"github.com/BiLuoHui/CommercialServiceSimple/pkg/protocol/rest"
)

func RunServer() error {
	// 创建Server
	ctx := context.Background()
	var port string
	flag.StringVar(&port, "http-port", "", "HTTP 端口号")
	flag.Parse()

	if len(port) == 0 {
		return fmt.Errorf("请指定HTTP端口号\n")
	}

	// 创建HTTP服务器
	return rest.RunServer(ctx, port)
}
