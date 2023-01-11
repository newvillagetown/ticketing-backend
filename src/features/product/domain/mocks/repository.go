package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"main/common/dbCommon/mysqlCommon"
	_interface "main/features/product/usecase/interface"
)

var _ _interface.IGetProductRepository = &Repository{}

type Repository struct {
	mock.Mock
}

func (r *Repository) FindOneProduct(ctx context.Context, productID string) (mysqlCommon.GormProduct, error) {
	ret := r.Called(ctx, productID)
	return ret.Get(0).(mysqlCommon.GormProduct), ret.Error(1)
}
