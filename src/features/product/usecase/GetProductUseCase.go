package usecase

import (
	"main/common/dbCommon/mysqlCommon"
	"main/features/product/model/request"
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

func (g *GetProductUseCase) Get(req request.ReqGetProduct) (mysqlCommon.Product, error) {
	productDTO, err := g.Repository.FindOneProduct(req.ProductID)
	if err != nil {
		return mysqlCommon.Product{}, err
	}

	return productDTO, nil
}
