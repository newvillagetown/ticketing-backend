package usecase

import (
	"context"
	"main/features/product/domain/request"
	"main/features/product/usecase/interface"
	"time"
)

type UpdateProductUseCase struct {
	Repository     _interface.IUpdateProductRepository
	ContextTimeout time.Duration
}

func NewUpdateProductUseCase(repo _interface.IUpdateProductRepository, timeout time.Duration) _interface.IUpdateProductUseCase {
	return &UpdateProductUseCase{
		Repository:     repo,
		ContextTimeout: timeout,
	}
}

func (u *UpdateProductUseCase) Update(c context.Context, req request.ReqUpdateProduct) error {
	ctx, cancel := context.WithTimeout(c, u.ContextTimeout)
	defer cancel()
	productDTO, err := u.Repository.FindOneProduct(ctx, req.ProductID)
	if err != nil {
		return err
	}
	updatedProductDTO := ConvertUpdateProductNewProductDTO(req, productDTO)
	err = u.Repository.FindOneAndUpdateProduct(ctx, updatedProductDTO)
	if err != nil {
		return err
	}
	return nil
}
