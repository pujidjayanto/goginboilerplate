package dto

type CreatePurchaseRequest struct {
	UserId    uint // todo: need to pass from auth token
	ProductId uint `json:"productId"`
	Quantity  int  `json:"quantity"` // todo: need gte validation
}
