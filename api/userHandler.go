package api

import (
	"encoding/json"
	"gohttp/dao"
	"gohttp/repo"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newItem dao.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	_, err := repo.DbConnection.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", newItem.Username, newItem.Email, newItem.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UsernameExist(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	rows, err := repo.DbConnection.Query(`SELECT * FROM users WHERE username = '` + params["username"] + `'`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	w.Header().Set("Content-Type", "application/json")
	if rows.Next() {
		json.NewEncoder(w).Encode("Exist")
	} else {
		json.NewEncoder(w).Encode("Available")
	}

}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var newItem dao.User
	var returnUser dao.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	rows, err := repo.DbConnection.Query(`SELECT * FROM users WHERE username = '` + newItem.Username + `' AND password = '` + newItem.Password + `'`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&returnUser.User_id, &returnUser.Username, &returnUser.Email, &returnUser.Password, &returnUser.Registration_date, &returnUser.Funds)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(returnUser)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}

}
