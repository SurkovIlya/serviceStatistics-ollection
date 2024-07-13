package server

import (
	"encoding/json"
	"mime"
	"net/http"

	"github.com/SurkovIlya/statistics-app/internal/model"
)

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
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	json.NewEncoder(w).Encode(res)
}

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
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	json.NewEncoder(w).Encode(http.StatusText(200))
}
