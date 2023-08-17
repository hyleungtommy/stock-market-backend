package api

import (
	"encoding/json"
	"gohttp/params"
	"gohttp/repo"
	"net/http"
)

func SellStock(w http.ResponseWriter, r *http.Request) {
	var param params.SellStock
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&param); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	_, err := repo.DbConnection.Exec("CALL sell_stock($1,$2,$3,$4)", param.UserId, param.StockId, param.SellAmt, param.SellPrice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func BuyStock(w http.ResponseWriter, r *http.Request) {
	var param params.BuyStock
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&param); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	_, err := repo.DbConnection.Exec("CALL buy_stock($1,$2)", param.UserId, param.TransactionId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func BuyInitStock(w http.ResponseWriter, r *http.Request) {
	var param params.BuyInitStock
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&param); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	_, err := repo.DbConnection.Exec("CALL buy_init_stock($1,$2,$3)", param.UserId, param.StockId, param.BuyAmt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
