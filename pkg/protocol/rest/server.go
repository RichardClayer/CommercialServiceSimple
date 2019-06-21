package rest

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "time"

    "github.com/BiLuoHui/CommercialServiceSimple/pkg/routes"
)

func RunServer(ctx context.Context, httpPort string) error {
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    mux := routes.NewRouter()

    srv := &http.Server{
        Addr:    ":" + httpPort,
        Handler: mux,
    }

    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)

    go func() {
        for range c {
            //
        }

        _, cancel := context.WithTimeout(ctx, 5*time.Second)
        defer cancel()
        _ = srv.Shutdown(ctx)
    }()

    log.Println("正在启动HTTP服务器……")

    return srv.ListenAndServe()
}
