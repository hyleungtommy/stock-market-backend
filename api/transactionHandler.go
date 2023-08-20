package api

import (
	"encoding/json"
	"gohttp/dao"
	"gohttp/params"
	"gohttp/repo"
	"net/http"
	"strconv"
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

func ListTransactions(w http.ResponseWriter, r *http.Request) {
	var filter params.TransFilter
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&filter); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	sql := `SELECT * FROM transactions `
	if filter.StockId > 0 {
		sql += ` WHERE stock_id = ` + strconv.Itoa(filter.StockId)
	}
	if filter.OrderbyRecent || filter.OrderbyAmt || filter.OrderbyPrice {
		sql += ` ORDER BY `
		if filter.OrderbyRecent {
			sql += ` post_date DESC `
			if filter.OrderbyPrice {
				sql += `, sell_price DESC `
			}
			if filter.OrderbyAmt {
				sql += `, sell_amt DESC `
			}
		} else if filter.OrderbyPrice {
			sql += ` sell_price DESC `
			if filter.OrderbyAmt {
				sql += `, sell_amt DESC `
			}
		} else if filter.OrderbyAmt {
			sql += ` sell_amt DESC `
		}
	}

	rows, err := repo.DbConnection.Query(sql)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []dao.Transaction
	for rows.Next() {
		var item dao.Transaction
		err := rows.Scan(&item.ID, &item.UserId, &item.StockId, &item.SellAmt, &item.SellPrice, &item.PostDate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
