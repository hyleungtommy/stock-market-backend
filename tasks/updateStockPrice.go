package tasks

import (
	"fmt"
	"gohttp/repo"
)

func UpdateStockPrice() {

	_, err := repo.DbConnection.Exec("CALL update_stock_price()")
	if err != nil {
		fmt.Println(`error on updating stock price:$1`, err)
		return
	}
}
