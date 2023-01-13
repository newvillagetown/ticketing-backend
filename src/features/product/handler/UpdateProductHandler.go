package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"main/common/jwtCommon"
	"main/common/valCommon"
	"main/features/product/domain/request"
	_interface "main/features/product/usecase/interface"
	"net/http"
)

type UpdateProductHandler struct {
	UseCase _interface.IUpdateProductUseCase
}

func NewUpdateProductHandler(c *echo.Echo, useCase _interface.IUpdateProductUseCase) {
	handler := &UpdateProductHandler{
		UseCase: useCase,
	}
	c.PUT("/v0.1/features/product", handler.update, middleware.JWTWithConfig(jwtCommon.JwtConfig))
}

// Product update
// @Router /v0.1/features/product [put]
// @Summary 상품 수정하기
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
// @Param json body request.ReqUpdateProduct true "json body"
// @Produce json
// @Success 200 {object} bool
// @Failure 400 {object} errorCommon.ResError
// @Failure 500 {object} errorCommon.ResError
// @Tags product
func (u *UpdateProductHandler) update(c echo.Context) error {
	req := &request.ReqUpdateProduct{}
	if err := valCommon.ValidateReq(c, req); err != nil {
		return err
	}
	ctx := c.Request().Context()

	err := u.UseCase.Update(ctx, *req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, true)
}
