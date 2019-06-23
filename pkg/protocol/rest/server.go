package rest

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/BiLuoHui/CommercialServiceSimple/pkg/routes"
)

type ServiceServer struct {
	db *sql.DB
}

func (s *ServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, fmt.Errorf("数据库连接失败")
	}

	return c, nil
}

// 创建server
func NewServer(db *sql.DB) *ServiceServer {
	return &ServiceServer{
		db: db,
	}
}

// 运行server
func RunServer(ctx context.Context, httpPort string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := routes.NewRouter()

	srv := &http.Server{
		Addr:    ":" + httpPort,
		Handler: mux,
	}

	go func() {
		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		_ = srv.Shutdown(ctx)
	}()

	log.Println("正在启动HTTP服务器……")

	return srv.ListenAndServe()
}
