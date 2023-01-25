package handler

import (
	"github.com/bxcodec/faker"
	"github.com/labstack/echo/v4"
	"main/common/dbCommon/mysqlCommon"
	"net/http"
	"net/http/httptest"
)

func (s *HandlerSuite) TestController_Update() {
	// Given
	var mockProducts []mysqlCommon.GormProduct
	err := faker.FakeData(&mockProducts)
	s.NoError(err)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/v0.1/features/product", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	c := s.engine.NewContext(req, rec)
	s.UpdateProductUseCase.On("Update", c.Request().Context(), mockProducts).Return(nil).Once()

	// when
	err = s.UpdateProductHandler.Update(c)

	// then
	s.NoError(err)
	s.Equal(c.Response().Status, rec.Code)
}
