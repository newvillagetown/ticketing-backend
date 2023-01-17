package usecase

import (
	"context"
	"main/common/dbCommon/mysqlCommon"
	"main/features/product/usecase/interface"
	"time"
)

type GetsProductUseCase struct {
	Repository     _interface.IGetsProductRepository
	ContextTimeout time.Duration
}

func NewGetsProductUseCase(repo _interface.IGetsProductRepository, timeout time.Duration) _interface.IGetsProductUseCase {
	return &GetsProductUseCase{
		Repository:     repo,
		ContextTimeout: timeout,
	}
}

func (g *GetsProductUseCase) Gets(c context.Context) ([]mysqlCommon.GormProduct, error) {
	ctx, cancel := context.WithTimeout(c, g.ContextTimeout)
	defer cancel()
	productList, err := g.Repository.FindProduct(ctx)
	if err != nil {
		return nil, err
	}
	return productList, nil
}
