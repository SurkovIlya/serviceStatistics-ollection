package server

import (
	"context"
	"net/http"
	"time"

	"github.com/SurkovIlya/statistics-app/internal/model"
)

type OrdersStorage interface {
	GetOrderBook(exchange_name, pair string) ([]*model.DepthOrder, error)
	SaveOrderBook(exchange_name, pair string, orderBook []*model.DepthOrder) error
	GetOrderHistory(client *model.Client) ([]*model.HistoryOrder, error)
	SaveOrder(client *model.Client, order *model.HistoryOrder) error
}

type Server struct {
	httpServer *http.Server
}

func New(port string) *Server {
	s := &Server{
		httpServer: &http.Server{
			Addr:           ":" + port,
			Handler:        initRoutes(),
			MaxHeaderBytes: 1 << 20, // 1 MB
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
	}

	return s
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func initRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /orderbook/get", GetOrderBook)
	mux.HandleFunc("POST /orderbook/save", SaveOrderBook)
	mux.HandleFunc("POST /orderhistory/get", GetOrderHistory)
	mux.HandleFunc("POST /orderhistory/save", SaveOrderHistory)

	return mux
}
