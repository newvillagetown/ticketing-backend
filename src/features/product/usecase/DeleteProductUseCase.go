package usecase

import (
	"main/features/product/model/request"
	_interface "main/features/product/usecase/interface"
)

type DeleteProductUseCase struct {
	Repository _interface.IDeleteProductRepository
}

func NewDeleteProductUseCase(repo _interface.IDeleteProductRepository) _interface.IDeleteProductUseCase {
	return &DeleteProductUseCase{
		Repository: repo,
	}
}

func (d *DeleteProductUseCase) Delete(req request.ReqDeleteProduct) error {
	err := d.Repository.FindOneAndDeleteUpdateProduct(req.ProductID)
	if err != nil {
		return err
	}
	return nil
}
