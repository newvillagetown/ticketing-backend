package usecase

import (
	_interface "main/features/product/usecase/interface"
)

type GetsProductUseCase struct {
	Repository _interface.IGetsProductRepository
}

func NewGetsProductUseCase(repo _interface.IGetsProductRepository) _interface.IGetsProductUseCase {
	return &GetsProductUseCase{
		Repository: repo,
	}
}

func (g *GetsProductUseCase) Gets() error {
	return nil
}
