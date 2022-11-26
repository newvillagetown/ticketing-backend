package usecase

import (
	_interface "main/features/product/usecase/interface"
)

type UpdateProductUseCase struct {
	Repository _interface.IUpdateProductRepository
}

func NewUpdateProductUseCase(repo _interface.IUpdateProductRepository) _interface.IUpdateProductUseCase {
	return &UpdateProductUseCase{
		Repository: repo,
	}
}

func (u *UpdateProductUseCase) Update() error {

	return nil
}
