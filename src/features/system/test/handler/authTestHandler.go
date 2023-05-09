package handler

import (
	"github.com/labstack/echo/v4"
	_interface "main/features/system/test/usecase/interface"
	"net/http"
)

type AuthTestIHandler struct {
	UseCase _interface.IAuthTestUseCase
}

func NewAuthTestHandler(c *echo.Echo, useCase _interface.IAuthTestUseCase) _interface.IAuthTestHandler {
	handler := &AuthTestIHandler{
		UseCase: useCase,
	}
	c.GET("/v0.1/test/auth", handler.AuthTest)
	return handler
}

// auth api test
// @Router /v0.1/test/auth [get]
// @Summary 인증 api 테스트
// @Description
// @Description ■ subscriptionType 구독 종류
// @Description BASIC(1) PLUS(2) B2B(3) B2H(4)
// @Description
// @Description ■ state 구독 상태
// @Description READY(1) WAIT(2) SUBSCRIBING(3) CANCEL(4) EXCHANGE(5) RESERVE_TERMINATE(6) RESERVE_PAYMENT(7)
// @Description
// @Description ■ provider 결제 수단
// @Description CARD(1) NAVER(2) KAKAO(3)
// @Description
// @Description ■ category 제품 유형
// @Description DEVICE(1) MOUTHPIECE(2)
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
// @Success 200 {object} bool
// @Failure 400 {object} errorSystem.ResError
// @Failure 401 {object} errorSystem.ResError
// @Failure 500 {object} errorSystem.ResError
// @Tags test
func (s *AuthTestIHandler) AuthTest(c echo.Context) error {

	ctx := c.Request().Context()
	err := s.UseCase.AuthTest(ctx)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, true)
}
