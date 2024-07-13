package orders

import (
	"fmt"

	"github.com/SurkovIlya/statistics-app/internal/model"
)

type PostgresStorage interface {
	SelectOrder(exchange_name, pair string) ([]*model.DepthOrder, error)
	InsertOrderBook(exchange_name, pair string, orderBook []*model.DepthOrder) error
	SelectHistory(client *model.Client) ([]*model.HistoryOrder, error)
	InsertHistory(client *model.Client, order *model.HistoryOrder) error
}

type OrderManager struct {
	QueryStorage PostgresStorage
}

func New(query PostgresStorage) *OrderManager {
	return &OrderManager{
		QueryStorage: query,
	}
}

func (om *OrderManager) GetOrderBook(exchange_name, pair string) ([]*model.DepthOrder, error) {
	res, err := om.QueryStorage.SelectOrder(exchange_name, pair)
	if err != nil {
		return nil, fmt.Errorf("error get order from BD: %s", err)
	}

	return res, nil
}
func (om *OrderManager) SaveOrderBook(exchange_name, pair string, orderBook []*model.DepthOrder) error {
	err := om.QueryStorage.InsertOrderBook(exchange_name, pair, orderBook)
	if err != nil {
		return fmt.Errorf("error save order in BD: %s", err)
	}

	return nil
}
func (om *OrderManager) GetOrderHistory(client *model.Client) ([]*model.HistoryOrder, error) {
	res, err := om.QueryStorage.SelectHistory(client)
	if err != nil {
		return nil, fmt.Errorf("error get history from BD: %s", err)
	}

	return res, nil
}
func (om *OrderManager) SaveOrder(client *model.Client, order *model.HistoryOrder) error {
	err := om.QueryStorage.InsertHistory(client, order)
	if err != nil {
		return fmt.Errorf("error save history in BD: %s", err)
	}

	return nil
}
