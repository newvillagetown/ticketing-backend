package handler

import (
	"github.com/labstack/echo/v4"
	"main/common/dbCommon/mongodbCommon"
	"main/features/product/repository"
	"main/features/product/usecase"
	_interface "main/features/product/usecase/interface"
	"net/http"
)

type GetProductHandler struct {
	UseCase _interface.IGetProductUseCase
}

func NewGetProductHandler() *GetProductHandler {
	return &GetProductHandler{UseCase: usecase.NewGetProductUseCase(repository.NewGetProductRepository(mongodbCommon.TokenCollection))}
}

// Product get
// @Router /v0.1/features/product [get]
// @Summary 상품 상세정보 가져오기
// @Description
// @Description ■ errCode with 500
// @Description INTERNAL_SERVER : 내부 로직 처리 실패
// @Description INTERNAL_DB : DB 처리 실패
// @Param token header string true "accessToken"
// @Param name formData string true "name"
// @Param description formData string true "description"
// @Param category formData string true "category"
// @Param image formData file false "image"
// @Param perAmount formData int true "perAmount"
// @Param totalCount formData int true "totalCount"
// @Param restCount formData int true "restCount"
// @Param startDate formData int true "startDate"
// @Param endDate formData int true "endDate"
// @Produce json
// @Success 200 {object} bool
// @Failure 400 {object} errorCommon.ResError
// @Failure 500 {object} errorCommon.ResError
// @Tags product
func (g *GetProductHandler) get(c echo.Context) error {

	return c.JSON(http.StatusCreated, true)
}
