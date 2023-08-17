package api

import (
	"encoding/json"
	"gohttp/dao"
	"gohttp/repo"
	"net/http"

	"github.com/gorilla/mux"
)

func ListStocks(w http.ResponseWriter, r *http.Request) {
	rows, err := repo.DbConnection.Query("SELECT * FROM stocks")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []dao.Stock
	for rows.Next() {
		var item dao.Stock
		err := rows.Scan(&item.ID, &item.Name, &item.Code, &item.InitPrice, &item.InitStock, &item.RemainStock, &item.CurrentPrice)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func GetStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	rows, err := repo.DbConnection.Query("SELECT * FROM stocks WHERE id = " + params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var item dao.Stock
	if rows.Next() {
		err := rows.Scan(&item.ID, &item.Name, &item.Code, &item.InitPrice, &item.InitStock, &item.RemainStock, &item.CurrentPrice)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func UpdateRemainStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var updatedItem dao.Stock
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	_, err := repo.DbConnection.Exec("UPDATE stocks SET remain_stock = $1 WHERE id = $2", updatedItem.RemainStock, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func UpdateCurrentPrice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var updatedItem dao.Stock
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	_, err := repo.DbConnection.Exec("UPDATE stocks SET current_price = $1 WHERE id = $2", updatedItem.CurrentPrice, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func ErrorResponse(w http.ResponseWriter, err error, statusCode int) {
	http.Error(w, err.Error(), statusCode)
}
