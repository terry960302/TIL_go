package model

type StoreInfo struct {
	Address   string      `json:"addr"`
	Code      interface{} `json:"code"` //string인게 있고 숫자인게 있어서 any로 설정
	Latitude  float64     `json:"lat"`
	Longitude float64     `json:"lng"`
	Name      string      `json:"name"`
	Type      string      `json:"type"`
}

type SortedStore struct {
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	Address     string  `json:"addr"`
	Type        string  `json:"type"`
	Latitude    float64 `json:"lat"`
	Longitude   float64 `json:"lng"`
	StockAt     string  `json:"stock_at"`
	RemainState string  `json:"remain_stat"`
	CreatedAt   string  `json:"created_at"`
}
