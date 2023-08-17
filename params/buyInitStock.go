package params

type BuyInitStock struct {
	UserId  int `json:"user_id"`
	StockId int `json:"stock_id"`
	BuyAmt  int `json:"buy_amt"`
}
