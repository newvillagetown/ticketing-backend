package handler

import (
	"github.com/labstack/echo/v4"
	"main/common/dbCommon/mongodbCommon"
	"main/common/valCommon"
	"main/features/product/model/request"
	"main/features/product/repository"
	"main/features/product/usecase"
	_interface "main/features/product/usecase/interface"
	"net/http"
)

type DeleteProductHandler struct {
	UseCase _interface.IDeleteProductUseCase
}

func NewDeleteProductHandler() *DeleteProductHandler {
	return &DeleteProductHandler{UseCase: usecase.NewDeleteProductUseCase(repository.NewDeleteProductRepository(mongodbCommon.TokenCollection))}
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
func (d *DeleteProductHandler) delete(c echo.Context) error {
	req := &request.ReqDeleteProduct{}
	if iErr := valCommon.ValidateReq(c, req); iErr != nil {
		return iErr
	}
	err := d.UseCase.Delete(*req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, true)
}
