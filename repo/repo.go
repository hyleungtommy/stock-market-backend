package repo

import (
	"database/sql"
	"gohttp/db"
)

var DbConnection *sql.DB

func SetupRepo() (err error) {
	DbConnection, err = db.GetPostgres()
	err = DbConnection.Ping()
	if err != nil {
		panic(err)
	}
	return
}

func CloseRepo() {
	DbConnection.Close()
}
