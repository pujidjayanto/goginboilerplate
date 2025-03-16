package dto

type CreatePurchaseRequest struct {
	UserId    uint
	ProductId uint `json:"productId" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,gt=0"`
}

// note: if need more validation use Validate() implements the binding.Validator interface
