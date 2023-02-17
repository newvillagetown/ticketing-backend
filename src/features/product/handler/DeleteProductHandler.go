package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"main/common/valCommon"
	"main/features/product/domain/request"
	"main/features/product/usecase/interface"
	"net/http"
)

type DeleteProductHandler struct {
	UseCase _interface.IDeleteProductUseCase
}

func NewDeleteProductHandler(c *echo.Echo, useCase _interface.IDeleteProductUseCase) _interface.IDeleteProductHandler {
	handler := &DeleteProductHandler{
		UseCase: useCase,
	}
	//	c.DELETE("/v0.1/features/product", handler.Delete, middleware.JWTWithConfig(jwtCommon.JwtConfig))
	c.DELETE("/v0.1/features/product", handler.Delete)
	return handler
}

// Product delete
// @Router /v0.1/features/product [delete]
// @Summary 상품 삭제하기(소프트)
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
// @Param json body request.ReqDeleteProduct true "json body"
// @Produce json
// @Success 200 {object} bool
// @Failure 400 {object} errorCommon.ResError
// @Failure 500 {object} errorCommon.ResError
// @Tags product
func (d *DeleteProductHandler) Delete(c echo.Context) error {
	req := &request.ReqDeleteProduct{}
	if err := valCommon.ValidateReq(c, req); err != nil {
		return err
	}
	fmt.Println(req.ProductID)
	ctx := c.Request().Context()
	err := d.UseCase.Delete(ctx, *req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, true)
}
