package _interface

import "main/common/dbCommon/mysqlCommon"

type IRegisterProductRepository interface {
	CreateProduct(productDTO mysqlCommon.Product) error
}

type IGetProductRepository interface {
	FindOneProduct() error
}
