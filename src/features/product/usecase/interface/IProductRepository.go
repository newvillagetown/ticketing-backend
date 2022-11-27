package _interface

import (
	"main/common/dbCommon/mysqlCommon"
)

type IRegisterProductRepository interface {
	CreateProduct(productDTO mysqlCommon.Product) error
}

type IGetProductRepository interface {
	FindOneProduct(productID string) (mysqlCommon.Product, error)
}

type IGetsProductRepository interface {
	FindProduct() ([]mysqlCommon.Product, error)
}

type IDeleteProductRepository interface {
	FindOneAndDeleteUpdateProduct(productID string) error
}
type IUpdateProductRepository interface {
	FindOneProduct(productID string) (mysqlCommon.Product, error)
	FindOneAndUpdateProduct(productDTO mysqlCommon.Product) error
}
