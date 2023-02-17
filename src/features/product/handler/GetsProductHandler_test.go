package handler

import (
	"github.com/bxcodec/faker"
	"main/common/dbCommon/mysqlCommon"
	"net/http"
	"net/http/httptest"
)

func (s *HandlerSuite) TestController_Gets() {
	// Given
	var mockProducts []mysqlCommon.GormProduct
	err := faker.FakeData(&mockProducts)
	s.NoError(err)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v0.1/features/product/gets", nil)
	c := s.engine.NewContext(req, rec)
	s.GetsProductUseCase.On("Gets", c.Request().Context()).Return(mockProducts, nil).Once()

	// when
	err = s.GetsProductHandler.Gets(c)

	// then
	s.NoError(err)
	s.Equal(c.Response().Status, rec.Code)
}
