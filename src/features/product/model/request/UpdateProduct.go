package request

type ReqUpdateProduct struct {
	ProductID string `json:"productID" validate:"required"`
}
