package handler

import (
	"encoding/json"
	"github.com/bxcodec/faker"
	"github.com/labstack/echo/v4"
	"main/common/dbCommon/mysqlCommon"
	"main/features/product/domain/request"
	"net/http"
	"net/http/httptest"
	"strings"
)

func (s *HandlerSuite) TestController_Delete() {
	// Given
	var mockProduct mysqlCommon.GormProduct
	err := faker.FakeData(&mockProduct)
	s.NoError(err)
	var reqDeleteProduct request.ReqDeleteProduct
	reqDeleteProduct.ProductID = mockProduct.GormModel.ID
	jsonBytes, err := json.Marshal(reqDeleteProduct)
	s.NoError(err)
	jsonString := string(jsonBytes)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/v0.1/features/product", strings.NewReader(jsonString))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	c := s.engine.NewContext(req, rec)
	s.DeleteProductUseCase.On("Delete", c.Request().Context(), reqDeleteProduct).Return(nil).Once()

	// when
	err = s.DeleteProductHandler.Delete(c)

	// then
	s.NoError(err)
	s.Equal(c.Response().Status, rec.Code)
}
