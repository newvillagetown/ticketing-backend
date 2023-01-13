package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"main/common/jwtCommon"
	"main/common/valCommon"
	"main/features/user/model/request"
	_interface "main/features/user/usecase/interface"
	"net/http"
)

type WithdrawalUserHandler struct {
	UseCase _interface.IWithdrawalUserUseCase
}

func NewWithdrawalUserHandler(c *echo.Echo, useCase _interface.IWithdrawalUserUseCase) {
	handler := &WithdrawalUserHandler{
		UseCase: useCase,
	}
	c.GET("/v0.1/auth/user/withdrawal", handler.WithdrawalUser, middleware.JWTWithConfig(jwtCommon.JwtConfig))
}

// withdrawal user
// @Router /v0.1/auth/user/withdrawal [post]
// @Summary 회원 탈퇴
// @Description
// @Description ■ errCode with 500
// @Description INTERNAL_SERVER : 내부 로직 처리 실패
// @Description INTERNAL_DB : DB 처리 실패
// @Produce json
// @Param json body request.ReqWithdrawalUser true "json body"
// @Success 200 {object} bool
// @Failure 400 {object} errorCommon.ResError
// @Failure 500 {object} errorCommon.ResError
// @Tags auth
func (w *WithdrawalUserHandler) WithdrawalUser(c echo.Context) error {
	req := &request.ReqWithdrawalUser{}
	if err := valCommon.ValidateReq(c, req); err != nil {
		return err
	}
	ctx := c.Request().Context()
	err := w.UseCase.WithdrawalUser(ctx, req.UserID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, true)
}
