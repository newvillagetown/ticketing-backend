package usecase

import (
	_interface "main/features/product/usecase/interface"
)

type GetProductUseCase struct {
	Repository _interface.IGetProductRepository
}

func NewGetProductUseCase(repo _interface.IGetProductRepository) _interface.IGetProductUseCase {
	return &GetProductUseCase{
		Repository: repo,
	}
}

func (g *GetProductUseCase) Get() error {

	return nil
}
