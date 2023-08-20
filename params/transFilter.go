package params

type TransFilter struct {
	StockId       int  `json:"stock_id"`
	OrderbyRecent bool `json:"order_by_recent"`
	OrderbyPrice  bool `json:"order_by_price"`
	OrderbyAmt    bool `json:"order_by_amt"`
}
