package rest

import (
	"context"
	"github.com/BiLuoHui/CommercialServiceSimple/pkg/routes"
	"log"
	"net/http"
)

// 运行server
func RunServer(ctx context.Context, httpPort string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := routes.NewRouter()

	srv := &http.Server{
		Addr:    ":" + httpPort,
		Handler: mux,
	}

	log.Println("正在启动HTTP服务器……")

	return srv.ListenAndServe()
}
