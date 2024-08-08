package server

import (
	"encoding/json"
	"mime"
	"net/http"

	"github.com/SurkovIlya/statistics-app/internal/model"
)

// @Summary Get order history
// @Tags orderhistory
// @Description get order history from DB
// @ID get-orderhistory
// @Accept  json
// @Produce  json
// @Param input body model.Client true "{"client_name": "John Doe", "exchange_name": "Example Exchange", "label": "Order123","pair": "BTC/USDT"}"
// @Success 200 {object} []model.HistoryOrder "[{"client_name": "John Doe","exchange_name": "Example Exchange","label": "Order123","pair": "BTC/USDT",	"side": "Buy","type": "Limit","base_qty": 1.5,"price": 40000.25,"algorithm_name_placed": "AlgorithmXYZ","lowest_sell_prc": 40200.75,"highest_buy_prc": 39950.5,"commission_quote_qty": 2,"time_placed": "2022-01-15T10:30:00Z"}]"
// @Router /orderhistory/get [post]
func (s *Server) GetOrderHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	contentType := r.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)

		return
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var rp *model.Client

	if err := dec.Decode(&rp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	res, err := s.orders.GetOrderHistory(rp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	json.NewEncoder(w).Encode(res)
}

// @Summary Save order history
// @Tags orderhistory
// @Description save order history in DB
// @ID save-orderhistory
// @Accept  json
// @Produce  json
// @Param input body model.HistoryOrder true "{"client_name": "John Doe","exchange_name": "Example Exchange","label": "Order123","pair": "BTC/USDT",	"side": "Buy","type": "Limit","base_qty": 1.5,"price": 40000.25,"algorithm_name_placed": "AlgorithmXYZ","lowest_sell_prc": 40200.75,"highest_buy_prc": 39950.5,"commission_quote_qty": 2,"time_placed": "2022-01-15T10:30:00Z"}"
// @Success 200 {string} string "OK"
// @Router /orderhistory/save [post]
func (s *Server) SaveOrderHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	contentType := r.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)

		return
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var rp *model.HistoryOrder

	if err := dec.Decode(&rp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	client := &model.Client{
		ClientName:   rp.ClientName,
		ExchangeName: rp.ExchangeName,
		Label:        rp.Label,
		Pair:         rp.Pair,
	}

	err = s.orders.SaveOrder(client, rp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	json.NewEncoder(w).Encode(http.StatusText(200))
}
