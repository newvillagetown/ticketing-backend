package _interface

import (
	"main/common/dbCommon/mysqlCommon"
	"main/features/product/model/request"
)

type IRegisterProductUseCase interface {
	Register(req request.ReqRegisterProduct) error
}

type IGetProductUseCase interface {
	Get(req request.ReqGetProduct) (mysqlCommon.GormProduct, error)
}

type IGetsProductUseCase interface {
	Gets() ([]mysqlCommon.GormProduct, error)
}

type IDeleteProductUseCase interface {
	Delete(req request.ReqDeleteProduct) error
}

type IUpdateProductUseCase interface {
	Update(req request.ReqUpdateProduct) error
}
