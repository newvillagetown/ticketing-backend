package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"main/common/dbCommon/mysqlCommon"
	"main/features/product/domain/request"
	_interface "main/features/product/usecase/interface"
)

var _ _interface.IGetProductUseCase = &Usecase{}

type Usecase struct {
	mock.Mock
}

func (u *Usecase) Get(c context.Context, req request.ReqGetProduct) (mysqlCommon.GormProduct, error) {
	ret := u.Called(c, req.ProductID)
	return ret.Get(0).(mysqlCommon.GormProduct), ret.Error(1)
}
