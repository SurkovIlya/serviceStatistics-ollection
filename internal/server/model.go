package server

import "github.com/SurkovIlya/statistics-app/internal/model"

type SaveOrderReq struct {
	ExchangeName string          `json:"exchange_name"`
	Pair         string          `json:"pair"`
	OrderBook    model.OrderBook `json:"order_book"`
}

type GetOrderReq struct {
	ExchangeName string `json:"exchange_name"`
	Pair         string `json:"pair"`
}
