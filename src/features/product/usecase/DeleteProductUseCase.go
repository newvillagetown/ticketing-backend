package usecase

import (
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

func (d *DeleteProductUseCase) Delete() error {
	return nil
}
