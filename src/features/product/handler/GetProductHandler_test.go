package handler

import (
	"context"
	"github.com/bxcodec/faker"
	"github.com/labstack/echo/v4"
	"main/common/dbCommon/mysqlCommon"
	"main/common/jwtCommon"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

func (s *HandlerSuite) TestController_Gets() {
	// Given
	var mockProducts []mysqlCommon.GormProduct
	err := faker.FakeData(&mockProducts)
	s.NoError(err)
	ctx := context.TODO()
	s.GetsProductUseCase.On("Gets", ctx).Return(mockProducts, nil).Once()

	req, err := http.NewRequestWithContext(context.TODO(), echo.GET, "/v0.1/features/product/gets", strings.NewReader(""))
	s.NoError(err)

	//쿠키 생성
	accessToken, _, err := jwtCommon.GenerateToken("ryan@breathings.co.kr", time.Now(), "259cb748-636e-4e06-9524-980a28127043")
	cookie := CreateCookie(accessToken)
	req.AddCookie(cookie)

	rec := httptest.NewRecorder()
	c := s.engine.NewContext(req, rec)
	handler := GetsProductHandler{UseCase: s.GetsProductUseCase}
	err = handler.Gets(c)
	s.NoError(err)

	// When
	res, err := s.GetsProductHandler.Gets(c)

	// Then
	s.NoError(err)
	s.Equal(http.StatusOK, rec.Code)
	s.Equal(mockProducts, res)
}
func CreateCookie(accessToken string) *http.Cookie {
	cookie := &http.Cookie{}
	cookie.Name = "accessToken"
	cookie.Value = accessToken
	cookie.Path = "/"
	cookie.SameSite = http.SameSiteLaxMode
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.Expires = time.Now().Add(1 * time.Hour)
	return cookie
}
