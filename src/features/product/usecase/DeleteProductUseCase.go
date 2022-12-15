package usecase

import (
	"context"
	"main/features/product/domain/request"
	_interface "main/features/product/usecase/interface"
	"time"
)

type DeleteProductUseCase struct {
	Repository     _interface.IDeleteProductRepository
	ContextTimeout time.Duration
}

func NewDeleteProductUseCase(repo _interface.IDeleteProductRepository, timeout time.Duration) _interface.IDeleteProductUseCase {
	return &DeleteProductUseCase{
		Repository:     repo,
		ContextTimeout: timeout,
	}
}

func (d *DeleteProductUseCase) Delete(c context.Context, req request.ReqDeleteProduct) error {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()
	err := d.Repository.FindOneAndDeleteUpdateProduct(ctx, req.ProductID)
	if err != nil {
		return err
	}
	return nil
}
