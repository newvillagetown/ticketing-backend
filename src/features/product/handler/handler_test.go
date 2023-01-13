package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"main/features/product/domain/mocks"
	"testing"
)

type HandlerSuite struct {
	suite.Suite
	engine *echo.Echo

	GetsProductHandler IGetsProductHandler
	GetsProductUseCase *mocks.GetsProductUseCase
}

func TestInit(t *testing.T) {
	suite.Run(t, new(HandlerSuite))
}

func (s *HandlerSuite) SetupSuite() {
	s.engine = echo.New()
	s.GetsProductUseCase = new(mocks.GetsProductUseCase)
}
