package usecase

import (
	"context"
	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v2"
	"main/features/product/domain/mocks"
	_interface "main/features/product/usecase/interface"
	"testing"
	"time"
)

type UseCaseSuite struct {
	suite.Suite
	mock sqlmock.Sqlmock
	repo *mocks.Repository

	DBTime            time.Duration
	ctx               context.Context
	GetProductUseCase _interface.IGetProductUseCase
}

func TestInit(t *testing.T) {
	suite.Run(t, new(UseCaseSuite))
}

func (s *UseCaseSuite) SetupSuite() {
	s.DBTime = 8 * time.Second
	s.repo = new(mocks.Repository)
	ctx, _ := context.WithTimeout(context.TODO(), s.DBTime)
	s.ctx = ctx
	s.GetProductUseCase = NewGetProductUseCase(s.repo, s.DBTime)
}
