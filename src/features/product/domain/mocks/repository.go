package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"main/common/dbCommon/mysqlCommon"
	_interface "main/features/product/usecase/interface"
)

var (
	_ _interface.IGetProductRepository      = &GetProductRepository{}
	_ _interface.IGetsProductRepository     = &GetsProductRepository{}
	_ _interface.IDeleteProductRepository   = &DeleteProductRepository{}
	_ _interface.IRegisterProductRepository = &RegisterProductRepository{}
)

type GetProductRepository struct {
	mock.Mock
}

type GetsProductRepository struct {
	mock.Mock
}

type DeleteProductRepository struct {
	mock.Mock
}
type RegisterProductRepository struct {
	mock.Mock
}
type UpdateProductRepository struct {
	mock.Mock
}

func (r *GetProductRepository) FindOneProduct(ctx context.Context, productID string) (mysqlCommon.GormProduct, error) {
	ret := r.Called(ctx, productID)
	return ret.Get(0).(mysqlCommon.GormProduct), ret.Error(1)
}

func (r *GetsProductRepository) FindProduct(ctx context.Context) ([]mysqlCommon.GormProduct, error) {
	ret := r.Called(ctx)
	return ret.Get(0).([]mysqlCommon.GormProduct), ret.Error(1)
}

func (r *DeleteProductRepository) FindOneAndDeleteUpdateProduct(ctx context.Context, productID string) error {
	ret := r.Called(ctx, productID)
	return ret.Error(0)
}

func (r *RegisterProductRepository) CreateProduct(ctx context.Context, productDTO mysqlCommon.GormProduct) error {
	ret := r.Called(ctx, productDTO)
	return ret.Error(0)
}

func (r *UpdateProductRepository) FindOneProduct(ctx context.Context, productID string) (mysqlCommon.GormProduct, error) {
	ret := r.Called(ctx, productID)
	return ret.Get(0).(mysqlCommon.GormProduct), ret.Error(1)
}

func (r *UpdateProductRepository) FindOneAndUpdateProduct(ctx context.Context, productDTO mysqlCommon.GormProduct) error {
	ret := r.Called(ctx, productDTO)
	return ret.Error(0)
}
