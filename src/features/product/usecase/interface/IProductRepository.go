package _interface

import (
	"context"
	"main/common/dbCommon/mysqlCommon"
)

type IRegisterProductRepository interface {
	CreateProduct(ctx context.Context, productDTO mysqlCommon.GormProduct) error
}

type IGetProductRepository interface {
	FindOneProduct(ctx context.Context, productID string) (mysqlCommon.GormProduct, error)
}

type IGetsProductRepository interface {
	FindProduct(ctx context.Context) ([]mysqlCommon.GormProduct, error)
}

type IDeleteProductRepository interface {
	FindOneAndDeleteUpdateProduct(ctx context.Context, productID string) error
}
type IUpdateProductRepository interface {
	FindOneProduct(ctx context.Context, productID string) (mysqlCommon.GormProduct, error)
	FindOneAndUpdateProduct(ctx context.Context, productDTO mysqlCommon.GormProduct) error
}
