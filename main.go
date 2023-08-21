package main

import (
	"fmt"
	"net/http"

	"gohttp/api"
	"gohttp/repo"
	"gohttp/tasks"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	go startServer()

	go startTasks()

	select {}

}

func startServer() {
	if err := repo.SetupRepo(); err != nil {
		panic(err)
	}

	defer repo.CloseRepo()
	router := mux.NewRouter()

	//for handling preflight CORS
	router.Methods("OPTIONS").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			return
		})

	router.HandleFunc("/stocks", api.ListStocks).Methods("GET")
	router.HandleFunc("/stocks/{id}", api.GetStock).Methods("GET")
	router.HandleFunc("/stocks/{id}/remain", api.UpdateRemainStock).Methods("PUT")
	router.HandleFunc("/stocks/{id}/price", api.UpdateCurrentPrice).Methods("PUT")

	router.HandleFunc("/users", api.CreateUser).Methods("POST")
	router.HandleFunc("/users/{username}/exist", api.UsernameExist).Methods("GET")
	router.HandleFunc("/users/login", api.LoginUser).Methods("POST")

	router.HandleFunc("/trans/buy", api.BuyStock).Methods("POST")
	router.HandleFunc("/trans/sell", api.SellStock).Methods("POST")
	router.HandleFunc("/trans/buyInit", api.BuyInitStock).Methods("POST")
	router.HandleFunc("/trans/list", api.ListTransactions).Methods("POST")

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", router)
}

func startTasks() {
	ticker := time.NewTicker(3600 * time.Second) // Schedule tasks every hour
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			tasks.UpdateStockPrice()
		}
	}
}

/*
func listItems(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM stocks")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []Stock
	for rows.Next() {
		var item Stock
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
*/
/*
func createItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	_, err := db.Exec("INSERT INTO items (name, price) VALUES ($1, $2)", newItem.Name, newItem.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedItem Item
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	_, err = db.Exec("UPDATE items SET name = $1, price = $2 WHERE id = $3", updatedItem.Name, updatedItem.Price, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("DELETE FROM items WHERE id = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
*/
