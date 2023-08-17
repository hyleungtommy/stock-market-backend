package params

type SellStock struct {
	UserId    int     `json:"user_id"`
	StockId   int     `json:"stock_id"`
	SellAmt   int     `json:"sell_amt"`
	SellPrice float32 `json:"sell_price"`
}
