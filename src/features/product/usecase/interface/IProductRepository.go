package _interface

import (
	"context"
	"main/common/dbCommon/mysqlCommon"
)

type IRegisterProductRepository interface {
	CreateProduct(productDTO mysqlCommon.GormProduct) error
}

type IGetProductRepository interface {
	FindOneProduct(ctx context.Context, productID string) (mysqlCommon.GormProduct, error)
}

type IGetsProductRepository interface {
	FindProduct() ([]mysqlCommon.GormProduct, error)
}

type IDeleteProductRepository interface {
	FindOneAndDeleteUpdateProduct(productID string) error
}
type IUpdateProductRepository interface {
	FindOneProduct(productID string) (mysqlCommon.GormProduct, error)
	FindOneAndUpdateProduct(productDTO mysqlCommon.GormProduct) error
}
