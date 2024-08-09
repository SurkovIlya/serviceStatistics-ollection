package server

import (
	"encoding/json"
	"mime"
	"net/http"

	_ "github.com/SurkovIlya/statistics-app/docs"
	"github.com/SurkovIlya/statistics-app/internal/model"
)

const maxOrdersInBook = 100

// @Summary Get order book
// @Tags orderbook
// @Description get orders from DB
// @ID get-orderbook
// @Accept  json
// @Produce  json
// @Param input body GetOrderReq true "name of the exchange or and the designation of the trading pair"
// @Success 200 {object} []model.DepthOrder "[{"price": 331.4,"base_qty": 3.66},{"price": 222.02,"base_qty": 5.66}]"
// @Router /orderbook/get [post]
func (s *Server) GetOrderBook(w http.ResponseWriter, r *http.Request) {
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

	var rp GetOrderReq

	if err := dec.Decode(&rp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	res, err := s.orders.GetOrderBook(rp.ExchangeName, rp.Pair)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	json.NewEncoder(w).Encode(res)
}

// @Summary Save order book
// @Tags orderbook
// @Description save orders in DB
// @ID save-orderbook
// @Accept  json
// @Produce  json
// @Param input body SaveOrderReq true "Save order book"
// @Success 200 {string} string "OK"
// @Router /orderbook/save [post]
func (s *Server) SaveOrderBook(w http.ResponseWriter, r *http.Request) {
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

	var rp SaveOrderReq

	if err := dec.Decode(&rp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	book := make([]*model.DepthOrder, 0, maxOrdersInBook)
	book = append(book, rp.OrderBook.Asks, rp.OrderBook.Bids)

	err = s.orders.SaveOrderBook(rp.ExchangeName, rp.Pair, book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	json.NewEncoder(w).Encode(http.StatusText(200))
}
