package _interface

import (
	"context"
	"main/common/dbCommon/mysqlCommon"
	"main/features/product/model/request"
)

type IRegisterProductUseCase interface {
	Register(req request.ReqRegisterProduct) error
}

type IGetProductUseCase interface {
	Get(c context.Context, req request.ReqGetProduct) (mysqlCommon.GormProduct, error)
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
