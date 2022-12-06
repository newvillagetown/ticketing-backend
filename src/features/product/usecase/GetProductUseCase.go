package usecase

import (
	"context"
	"main/common/dbCommon/mysqlCommon"
	"main/features/product/model/request"
	_interface "main/features/product/usecase/interface"
	"time"
)

type GetProductUseCase struct {
	Repository     _interface.IGetProductRepository
	ContextTimeout time.Duration
}

func NewGetProductUseCase(repo _interface.IGetProductRepository, timeout time.Duration) _interface.IGetProductUseCase {
	return &GetProductUseCase{
		Repository:     repo,
		ContextTimeout: timeout,
	}
}

func (g *GetProductUseCase) Get(c context.Context, req request.ReqGetProduct) (mysqlCommon.GormProduct, error) {
	ctx, cancel := context.WithTimeout(c, g.ContextTimeout)
	defer cancel()
	productDTO, err := g.Repository.FindOneProduct(ctx, req.ProductID)
	if err != nil {
		return mysqlCommon.GormProduct{}, err
	}

	return productDTO, nil
}
