package orders

import (
	"github.com/SurkovIlya/statistics-app/internal/model"
	"github.com/SurkovIlya/statistics-app/pkg/postgres"
)

type OrderManager struct {
	storage *postgres.Database
}

func New(storage *postgres.Database) *OrderManager {
	return &OrderManager{
		storage: storage,
	}
}

func (om *OrderManager) GetOrderBook(exchange_name, pair string) ([]*model.DepthOrder, error) {
	return nil, nil
}
func (om *OrderManager) SaveOrderBook(exchange_name, pair string, orderBook []*model.DepthOrder) error {
	return nil
}
func (om *OrderManager) GetOrderHistory(client *model.Client) ([]*model.HistoryOrder, error) {
	return nil, nil
}
func (om *OrderManager) SaveOrder(client *model.Client, order *model.HistoryOrder) error {
	return nil
}
