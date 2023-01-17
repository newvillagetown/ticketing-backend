package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"main/common/awsCommon"
	"main/features/product/domain/mocks"
	_ "main/features/product/usecase/interface"
	_interface "main/features/product/usecase/interface"
	mw "main/middleware"
	"net/http"
	"net/http/httptest"
	"testing"
)

type HandlerSuite struct {
	suite.Suite
	engine *echo.Echo

	GetsProductUseCase *mocks.GetsProductUseCase

	GetsProductHandler _interface.IGetsProductHandler
}

// Write test definition with TestSuite.
func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(HandlerSuite))
}

func (s *HandlerSuite) SetupTest() {
	s.engine = echo.New()
	s.GetsProductUseCase = new(mocks.GetsProductUseCase)

	NewGetsProductHandler(s.engine, s.GetsProductUseCase)

	awsCommon.InitAws()
	mw.InitMiddleware(s.engine)

}

func (s *HandlerSuite) buildContextAndRecorder(httpRequest *http.Request) (ctx echo.Context, rec *httptest.ResponseRecorder) {
	rec = httptest.NewRecorder()
	ctx = s.engine.NewContext(httpRequest, rec)
	return
}
