package handler

import (
	"github.com/labstack/echo/v4"
	"main/common/dbCommon/mongodbCommon"
	"main/features/product/repository"
	"main/features/product/usecase"
	_interface "main/features/product/usecase/interface"
	"net/http"
)

type RegisterProductHandler struct {
	UseCase _interface.IRegisterProductUseCase
}

func NewRegisterProductHandler() *RegisterProductHandler {
	return &RegisterProductHandler{UseCase: usecase.NewRegisterProductUseCase(repository.NewRegisterProductRepository(mongodbCommon.TokenCollection))}
}

// Product Register
// @Router /v0.1/feature/product [post]
// @Summary 상품 등록
// @Description
// @Description ■ errCode with 500
// @Description INTERNAL_SERVER : 내부 로직 처리 실패
// @Description INTERNAL_DB : DB 처리 실패
// @Produce json
// @Success 200 {object} bool
// @Failure 400 {object} errorCommon.ResError
// @Failure 500 {object} errorCommon.ResError
// @Tags auth
func (r *RegisterProductHandler) post(c echo.Context) error {

	return c.JSON(http.StatusOK, true)
}
