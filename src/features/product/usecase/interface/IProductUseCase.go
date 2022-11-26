package _interface

import (
	"main/common/dbCommon/mysqlCommon"
	"main/features/product/model/request"
)

type IRegisterProductUseCase interface {
	Register(req request.ReqRegisterProduct) error
}

type IGetProductUseCase interface {
	Get(req request.ReqGetProduct) (mysqlCommon.Product, error)
}

type IGetsProductUseCase interface {
	Gets() ([]mysqlCommon.Product, error)
}

type IDeleteProductUseCase interface {
	Delete(req request.ReqDeleteProduct) error
}

type IUpdateProductUseCase interface {
	Update() error
}
