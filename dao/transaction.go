package dao

import "time"

type Transaction struct {
	ID        int       `json:"id"`
	UserId    int       `json:"user_id"`
	StockId   int       `json:"stock_id"`
	SellAmt   int       `json:"sell_amt"`
	SellPrice int       `json:"sell_price"`
	PostDate  time.Time `json:"post_date"`
}
