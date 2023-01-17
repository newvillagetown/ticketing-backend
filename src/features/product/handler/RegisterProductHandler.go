package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"main/common/jwtCommon"
	"main/features/product/domain/request"
	"main/features/product/usecase/interface"
	"net/http"
	"strconv"
)

type RegisterProductHandler struct {
	UseCase _interface.IRegisterProductUseCase
}

func NewRegisterProductHandler(c *echo.Echo, useCase _interface.IRegisterProductUseCase) {
	handler := &RegisterProductHandler{
		UseCase: useCase,
	}
	c.POST("/v0.1/features/product", handler.post, middleware.JWTWithConfig(jwtCommon.JwtConfig))
}

// Product Register
// @Router /v0.1/features/product [post]
// @Summary 상품 등록
// @Description
// @Description ■ errCode with 500
// @Description INTERNAL_SERVER : 내부 로직 처리 실패
// @Description INTERNAL_DB : DB 처리 실패
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
func (r *RegisterProductHandler) post(c echo.Context) error {

	name := c.FormValue("name")
	description := c.FormValue("description")
	category := c.FormValue("category")
	perAmount, err := strconv.Atoi(c.FormValue("perAmount"))
	if err != nil {
		return err
	}
	totalCount, err := strconv.Atoi(c.FormValue("totalCount"))
	if err != nil {
		return err
	}
	restCount, err := strconv.Atoi(c.FormValue("restCount"))
	if err != nil {
		return err
	}
	startDate, err := strconv.Atoi(c.FormValue("startDate"))
	if err != nil {
		return err
	}
	endDate, err := strconv.Atoi(c.FormValue("endDate"))
	if err != nil {
		return err
	}
	img, err := c.FormFile("image")
	if err != nil && img != nil {
		return err
	}
	req := request.ReqRegisterProduct{
		Name:        name,
		Description: description,
		Category:    category,
		PerAmount:   int64(perAmount),
		TotalCount:  int64(totalCount),
		RestCount:   int64(restCount),
		Image:       img,
		StartDate:   int64(startDate),
		EndDate:     int64(endDate),
	}
	ctx := c.Request().Context()

	err = r.UseCase.Register(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, true)
}
