package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"main/common/awsCommon"
	"main/features/product/domain/mocks"
	_ "main/features/product/usecase/interface"
	_interface "main/features/product/usecase/interface"
	mw "main/middleware"
	"testing"
)

type HandlerSuite struct {
	suite.Suite
	engine *echo.Echo

	GetsProductUseCase   *mocks.GetsProductUseCase
	GetProductUseCase    *mocks.GetProductUseCase
	DeleteProductUseCase *mocks.DeleteProductUseCase

	GetsProductHandler   _interface.IGetsProductHandler
	GetProductHandler    _interface.IGetProductHandler
	DeleteProductHandler _interface.IDeleteProductHandler
}

// Write test definition with TestSuite.
func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(HandlerSuite))
}

func (s *HandlerSuite) SetupTest() {
	s.engine = echo.New()
	s.GetsProductUseCase = new(mocks.GetsProductUseCase)
	s.GetProductUseCase = new(mocks.GetProductUseCase)
	s.DeleteProductUseCase = new(mocks.DeleteProductUseCase)

	s.GetsProductHandler = NewGetsProductHandler(s.engine, s.GetsProductUseCase)
	s.GetProductHandler = NewGetProductHandler(s.engine, s.GetProductUseCase)
	s.DeleteProductHandler = NewDeleteProductHandler(s.engine, s.DeleteProductUseCase)

	awsCommon.InitAws()
	mw.InitMiddleware(s.engine)

}
