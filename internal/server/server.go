package server

import (
	"context"
	"net/http"
	"time"

	"github.com/SurkovIlya/statistics-app/internal/model"
	httpSwagger "github.com/swaggo/http-swagger"
)

type OrdersStorage interface {
	GetOrderBook(exchange_name, pair string) ([]*model.DepthOrder, error)
	SaveOrderBook(exchange_name, pair string, orderBook []*model.DepthOrder) error
	GetOrderHistory(client *model.Client) ([]*model.HistoryOrder, error)
	SaveOrder(client *model.Client, order *model.HistoryOrder) error
}

type Server struct {
	httpServer *http.Server
	orders     OrdersStorage
}

func New(port string, orders OrdersStorage) *Server {
	s := &Server{
		httpServer: &http.Server{
			Addr:           ":" + port,
			MaxHeaderBytes: 1 << 20, // 1 MB
			ReadTimeout:    200 * time.Millisecond,
			WriteTimeout:   200 * time.Millisecond,
		},
		orders: orders,
	}

	s.initRoutes()

	return s
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) initRoutes() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /orderbook/get", s.GetOrderBook)
	mux.HandleFunc("POST /orderbook/save", s.SaveOrderBook)
	mux.HandleFunc("POST /orderhistory/get", s.GetOrderHistory)
	mux.HandleFunc("POST /orderhistory/save", s.SaveOrderHistory)

	mux.HandleFunc("/swagger/*", httpSwagger.WrapHandler)
	// mux.HandleFunc("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/swagger/doc.json")))

	s.httpServer.Handler = mux
}
