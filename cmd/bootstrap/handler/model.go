package handler

// swagger:model ProductRequest
type ProductRequest struct {
	// the id of basket
	BasketID string `json:"basket_id" binding:"required" example:"0bfce8da-bdc9-11ec-b9f3-acde48001122"`
	// the code of product
	ProductCode string `json:"product_code" binding:"required"`
}

// swagger:model BasketRequest
type BasketRequest struct {
	// the id for add a new basket
	BasketID string `json:"basket_id" binding:"required"`
}

// swagger:model Response
type Response struct {
	// basket id
	ID string `json:"basket_id"`
	// items
	Item []Item `json:"items"`
	// total
	Total float64 `json:"total"`
}

// swagger:model Product
type Product struct {
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// swagger:model Item
type Item struct {
	Product  Product `json:"product"`
	Quantity int     `json:"quantity"`
	Total    float64 `json:"total"`
}
