package dao

type Stock struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Code         string  `json:"code"`
	InitPrice    float32 `json:"init_price"`
	InitStock    int     `json:"init_stock"`
	RemainStock  int     `json:"remain_stock"`
	CurrentPrice float32 `json:"current_price"`
}
