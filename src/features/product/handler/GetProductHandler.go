package handler

import (
	"github.com/labstack/echo/v4"
	"main/common/dbCommon/mongodbCommon"
	"main/common/dbCommon/mysqlCommon"
	"main/common/valCommon"
	"main/features/product/model/request"
	"main/features/product/repository"
	"main/features/product/usecase"
	_interface "main/features/product/usecase/interface"
	"net/http"
)

type GetProductHandler struct {
	UseCase _interface.IGetProductUseCase
}

func NewGetProductHandler() *GetProductHandler {
	return &GetProductHandler{UseCase: usecase.NewGetProductUseCase(repository.NewGetProductRepository(mysqlCommon.GormDB, mongodbCommon.TokenCollection), mysqlCommon.DBTimeOut)}
}

// Product get
// @Router /v0.1/features/product [get]
// @Summary 상품 상세정보 가져오기
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
// @Param productID query string true "productID"
// @Produce json
// @Success 200 {object} response.ResGetProduct
// @Failure 400 {object} errorCommon.ResError
// @Failure 500 {object} errorCommon.ResError
// @Tags product
func (g *GetProductHandler) Get(c echo.Context) error {
	req := &request.ReqGetProduct{}
	if iErr := valCommon.ValidateReq(c, req); iErr != nil {
		return iErr
	}
	ctx := c.Request().Context()
	productDTO, err := g.UseCase.Get(ctx, *req)
	if err != nil {
		return err
	}
	res := usecase.ConvertGetProductToRes(productDTO)
	return c.JSON(http.StatusOK, res)
}
