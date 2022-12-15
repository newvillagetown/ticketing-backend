package _interface

import (
	"context"
	"main/common/dbCommon/mysqlCommon"
	"main/features/product/domain/request"
)

type IRegisterProductUseCase interface {
	Register(c context.Context, req request.ReqRegisterProduct) error
}

type IGetProductUseCase interface {
	Get(c context.Context, req request.ReqGetProduct) (mysqlCommon.GormProduct, error)
}

type IGetsProductUseCase interface {
	Gets(c context.Context) ([]mysqlCommon.GormProduct, error)
}

type IDeleteProductUseCase interface {
	Delete(c context.Context, req request.ReqDeleteProduct) error
}

type IUpdateProductUseCase interface {
	Update(c context.Context, req request.ReqUpdateProduct) error
}
