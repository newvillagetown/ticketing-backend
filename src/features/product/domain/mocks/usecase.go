package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"main/common/dbCommon/mysqlCommon"
	"main/features/product/domain/request"
	_interface "main/features/product/usecase/interface"
)

var (
	_ _interface.IGetProductUseCase    = &GetProductUseCase{}
	_ _interface.IGetsProductUseCase   = &GetsProductUseCase{}
	_ _interface.IDeleteProductUseCase = &DeleteProductUseCase{}
)

type GetProductUseCase struct {
	mock.Mock
}
type GetsProductUseCase struct {
	mock.Mock
}
type DeleteProductUseCase struct {
	mock.Mock
}

func (u *GetProductUseCase) Get(c context.Context, req request.ReqGetProduct) (mysqlCommon.GormProduct, error) {
	ret := u.Called(c, req.ProductID)
	return ret.Get(0).(mysqlCommon.GormProduct), ret.Error(1)
}

func (u *GetsProductUseCase) Gets(c context.Context) ([]mysqlCommon.GormProduct, error) {
	ret := u.Called(c)
	return ret.Get(0).([]mysqlCommon.GormProduct), ret.Error(1)
}

func (u *DeleteProductUseCase) Delete(c context.Context, req request.ReqDeleteProduct) error {
	ret := u.Called(c, req)
	return ret.Error(0)
}
