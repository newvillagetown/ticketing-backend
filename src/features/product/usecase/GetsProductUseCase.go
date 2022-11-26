package usecase

import (
	"main/common/dbCommon/mysqlCommon"
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

func (g *GetsProductUseCase) Gets() ([]mysqlCommon.Product, error) {
	productList, err := g.Repository.FindProduct()
	if err != nil {
		return nil, err
	}
	return productList, nil
}
