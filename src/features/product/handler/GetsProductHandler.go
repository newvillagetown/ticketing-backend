package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"main/common/jwtCommon"
	"main/features/product/usecase"
	_interface "main/features/product/usecase/interface"
	"net/http"
)

type GetsProductHandler struct {
	UseCase _interface.IGetsProductUseCase
}

func NewGetsProductHandler(c *echo.Echo, useCase _interface.IGetsProductUseCase) {
	handler := &GetsProductHandler{
		UseCase: useCase,
	}
	c.GET("/v0.1/features/product/gets", handler.Gets, middleware.JWTWithConfig(jwtCommon.JwtConfig))
}

// Product gets
// @Router /v0.1/features/product/gets [get]
// @Summary 상품 목록 가져오기
// @Description
// @Description ■ errCode with 400
// @Description PARAM_BAD : 파라미터 오류
// @Description
// @Description ■ errCode with 401
// @Description TOKEN_BAD : 토큰 인증 실패
// @Description POLICY_VIOLATION : 토큰 세션 정책 위반
// @Description
// @Description ■ errCode with 500
// @Description INTERNAL_SERVER : 내부 로직 처리 실패
// @Description INTERNAL_DB : DB 처리 실패
// @Produce json
// @Success 200 {object} response.ResGetsProduct
// @Failure 400 {object} errorCommon.ResError
// @Failure 500 {object} errorCommon.ResError
// @Tags product
func (g *GetsProductHandler) Gets(c echo.Context) error {
	ctx := c.Request().Context()
	productList, err := g.UseCase.Gets(ctx)
	if err != nil {
		return err
	}
	res := usecase.ConvertGetsProductToRes(productList)

	return c.JSON(http.StatusOK, res)
}
