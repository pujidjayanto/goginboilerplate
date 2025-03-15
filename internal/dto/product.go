package dto

type ProductItem struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Price    string `json:"price"`
	Quantity int    `json:"quantity"`
}

type GetAllProduct struct {
	Products []*ProductItem `json:"products"`
}
