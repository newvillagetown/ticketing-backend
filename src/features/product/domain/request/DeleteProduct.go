package request

type ReqDeleteProduct struct {
	ProductID string `json:"productID" validate:"required"`
}
