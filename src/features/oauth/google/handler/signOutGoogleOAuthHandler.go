package handler

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"main/common/dbCommon/mongodbCommon"
	"main/features/oauth/google/repository"
	"main/features/oauth/google/usecase"
	_interface "main/features/oauth/google/usecase/interface"
	"net/http"
)

type SignOutGoogleOAuthHandler struct {
	UseCase _interface.ISignOutGoogleOAuthUseCase
}

func NewSignOutGoogleOAuthHandler() *SignOutGoogleOAuthHandler {
	return &SignOutGoogleOAuthHandler{UseCase: usecase.NewSignOutGoogleOAuthUseCase(repository.NewSignOutGoogleOAuthRepository(mongodbCommon.TokenCollection))}
}

// GoogleSignOut
// @Router /v0.1/auth/google/signout [get]
// @Summary google 로그아웃
// @Description
// @Description ■ errCode with 500
// @Description INTERNAL_SERVER : 내부 로직 처리 실패
// @Description INTERNAL_DB : DB 처리 실패
// @Param token header string true "accessToken"
// @Produce json
// @Success 200 {object} bool
// @Failure 400 {object} errorCommon.ResError
// @Failure 500 {object} errorCommon.ResError
// @Tags auth
func (s *SignOutGoogleOAuthHandler) SignOutGoogle(c echo.Context) error {
	//jwt 파싱해서 크레임 정보 가져온다.
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	fmt.Println(claims["userID"].(string))
	err := s.UseCase.SignOutGoogle(email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, true)
}
