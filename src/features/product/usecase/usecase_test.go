package usecase

import (
	"context"
	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v2"
	"main/features/product/domain/mocks"
	"main/features/product/usecase/interface"
	"testing"
	"time"
)

type UseCaseSuite struct {
	suite.Suite
	mock   sqlmock.Sqlmock
	DBTime time.Duration
	ctx    context.Context

	GetProductRepo      *mocks.GetProductRepository
	GetsProductRepo     *mocks.GetsProductRepository
	DeleteProductRepo   *mocks.DeleteProductRepository
	RegisterProductRepo *mocks.RegisterProductRepository
	UpdateProductRepo   *mocks.UpdateProductRepository

	GetProductUseCase      _interface.IGetProductUseCase
	GetsProductUseCase     _interface.IGetsProductUseCase
	DeleteProductUseCase   _interface.IDeleteProductUseCase
	RegisterProductUseCase _interface.IRegisterProductUseCase
	UpdateProductUseCase   _interface.IUpdateProductUseCase
}

func TestInit(t *testing.T) {
	suite.Run(t, new(UseCaseSuite))
}

func (s *UseCaseSuite) SetupSuite() {
	s.DBTime = 8 * time.Second
	ctx, _ := context.WithTimeout(context.TODO(), s.DBTime)
	s.ctx = ctx

	s.GetProductRepo = new(mocks.GetProductRepository)
	s.GetsProductRepo = new(mocks.GetsProductRepository)
	s.DeleteProductRepo = new(mocks.DeleteProductRepository)
	s.RegisterProductRepo = new(mocks.RegisterProductRepository)
	s.UpdateProductRepo = new(mocks.UpdateProductRepository)

	s.GetProductUseCase = NewGetProductUseCase(s.GetProductRepo, s.DBTime)
	s.GetsProductUseCase = NewGetsProductUseCase(s.GetsProductRepo, s.DBTime)
	s.DeleteProductUseCase = NewDeleteProductUseCase(s.DeleteProductRepo, s.DBTime)
	s.RegisterProductUseCase = NewRegisterProductUseCase(s.RegisterProductRepo, s.DBTime)
	s.UpdateProductUseCase = NewUpdateProductUseCase(s.UpdateProductRepo, s.DBTime)
}
