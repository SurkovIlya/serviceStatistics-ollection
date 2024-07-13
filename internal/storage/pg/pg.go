package pg

import (
	"encoding/json"
	"fmt"

	"github.com/SurkovIlya/statistics-app/internal/model"
	"github.com/SurkovIlya/statistics-app/pkg/postgres"
)

const maxOrderScan = 1000

type PostgresStorage struct {
	storage *postgres.Database
}

func New(storage *postgres.Database) *PostgresStorage {
	return &PostgresStorage{
		storage: storage,
	}
}

func (psq *PostgresStorage) SelectOrder(exchange_name, pair string) ([]*model.DepthOrder, error) {
	query := `SELECT asks, bids FROM order_book WHERE exchange = $1 AND pair = $2`
	rows, err := psq.storage.Conn.Query(query, exchange_name, pair)
	if err != nil {
		return nil, fmt.Errorf("error query: %s", err)
	}
	defer rows.Close()

	res := make([]*model.DepthOrder, 0, maxOrderScan)

	for rows.Next() {
		var asksJSON interface{}
		var bidsJSON interface{}
		var asks *model.DepthOrder
		var bids *model.DepthOrder

		err := rows.Scan(&asksJSON, &bidsJSON)
		if err != nil {
			return nil, fmt.Errorf("error scan: %s", err)
		}

		err = json.Unmarshal(asksJSON.([]byte), &asks)
		if err != nil {
			return nil, fmt.Errorf("error Unmarshal asksJson: %s", err)
		}
		err = json.Unmarshal(bidsJSON.([]byte), &bids)
		if err != nil {
			return nil, fmt.Errorf("error Unmarshal bidsJson: %s", err)
		}

		res = append(res, asks, bids)
	}

	return res, nil
}

func (psq *PostgresStorage) InsertOrderBook(exchange_name, pair string, orderBook []*model.DepthOrder) error {
	query := `INSERT INTO order_book (exchange, pair, asks, bids) VALUES ($1, $2, $3, $4)`

	asksJson, err := json.Marshal(orderBook[0])
	if err != nil {
		return fmt.Errorf("error Marshal asksJson: %s", err)
	}
	bidsJson, err := json.Marshal(orderBook[1])
	if err != nil {
		return fmt.Errorf("error Marshal bidsJson: %s", err)
	}

	_, err = psq.storage.Conn.Exec(query, exchange_name, pair, asksJson, bidsJson)
	if err != nil {
		return fmt.Errorf("error insert: %s", err)
	}

	return nil
}

func (psq *PostgresStorage) SelectHistory(client *model.Client) ([]*model.HistoryOrder, error) {
	var res []*model.HistoryOrder

	query := `SELECT * FROM order_history 
					WHERE client_name = $1 AND exchange_name = $2 AND label = $3 AND pair = $4`
	rows, err := psq.storage.Conn.Query(query, client.ClientName, client.ExchangeName, client.Label, client.Pair)
	if err != nil {
		return nil, fmt.Errorf("error query: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var history model.HistoryOrder

		err := rows.Scan(&history.ClientName,
			&history.ExchangeName,
			&history.Label,
			&history.Pair,
			&history.Side,
			&history.Type,
			&history.BaseQty,
			&history.Price,
			&history.AlgorithmNamePlaced,
			&history.LowestSellPrc,
			&history.HighestBuyPrc,
			&history.CommissionQuoteQty,
			&history.TimePlaced)
		if err != nil {
			return nil, fmt.Errorf("error scan: %s", err)
		}

		res = append(res, &history)
	}

	return res, nil
}

func (psq *PostgresStorage) InsertHistory(client *model.Client, order *model.HistoryOrder) error {
	query := `INSERT INTO order_history (client_name, 
										exchange_name,
										label, 
										pair, 
										side, 
										type, 
										base_qty, 
										price, 
										algorithm_name_placed, 
										lowest_sell_prc, 
										highest_buy_prc, 
										commission_quote_qty, 
										time_placed) 
									VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`

	_, err := psq.storage.Conn.Exec(query, client.ClientName, client.ExchangeName, client.Label, client.Pair, order.Side, order.Type, order.BaseQty, order.Price, order.AlgorithmNamePlaced, order.LowestSellPrc, order.HighestBuyPrc, order.CommissionQuoteQty, order.TimePlaced)
	if err != nil {
		return fmt.Errorf("error insert order in table: %s", err)
	}

	return nil
}
