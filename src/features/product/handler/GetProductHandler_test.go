package handler

import (
	"github.com/bxcodec/faker"
	"github.com/labstack/echo/v4"
	"main/common/dbCommon/mysqlCommon"
	"main/features/product/domain/request"
	"net/http"
	"net/http/httptest"
	"net/url"
)

func (s *HandlerSuite) TestController_Get() {
	// Given
	var mockProduct mysqlCommon.GormProduct
	err := faker.FakeData(&mockProduct)
	s.NoError(err)
	var reqGetProduct request.ReqGetProduct
	reqGetProduct.ProductID = mockProduct.GormModel.ID

	rec := httptest.NewRecorder()
	q := make(url.Values)
	q.Set("productID", mockProduct.GormModel.ID)

	req := httptest.NewRequest(http.MethodGet, "/v0.1/features/product?"+q.Encode(), nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	c := s.engine.NewContext(req, rec)
	s.GetProductUseCase.On("Get", c.Request().Context(), reqGetProduct.ProductID).Return(mockProduct, nil).Once()

	// when
	err = s.GetProductHandler.Get(c)

	// then
	s.NoError(err)
	s.Equal(c.Response().Status, rec.Code)
}
