package handler

import (
	"encoding/base64"
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"main/common/errorCommon"
	"main/features/message/domain/request"
	_interface "main/features/message/usecase/interface"
	"net/http"
	"strings"
)

type SendNaverSmsHandler struct {
	UseCase _interface.ISendNaverSmsUseCase
}

func NewSendNaverSmsHandler(c *echo.Echo, useCase _interface.ISendNaverSmsUseCase) _interface.ISendNaverSmsHandler {
	handler := &SendNaverSmsHandler{
		UseCase: useCase,
	}
	//c.GET("/v0.1/features/product", handler.Get, middleware.JWTWithConfig(jwtCommon.JwtConfig))
	c.POST("/v0.1/features/message/naver/sms", handler.Send)
	return handler
}

// naver sms send
// @Router /v0.1/features/message/naver/sms [post]
// @Summary 네이버 sms 메시지 전송
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
// @Param phoneList formData []string true "phoneList"
// @Param contentType formData string true "contentType"
// @Param content formData string true "content"
// @Param smsType formData string true "smsType"
// @Param reserveTime formData string false "reserveTime"
// @Param scheduleCode formData string false "scheduleCode"
// @Param image formData file false "image"
// @Produce json
// @Success 200 {object} bool
// @Failure 400 {object} errorCommon.ResError
// @Failure 500 {object} errorCommon.ResError
// @Tags message
func (s *SendNaverSmsHandler) Send(c echo.Context) error {

	phoneList := c.FormValue("phoneList")
	contentType := c.FormValue("contentType")
	content := c.FormValue("content")
	smsType := c.FormValue("smsType")
	reserveTime := c.FormValue("reserveTime")
	scheduleCode := c.FormValue("scheduleCode")
	req := request.ReqSendNaverSms{}
	if phoneList == "" || contentType == "" || content == "" || smsType == "" {
		return fmt.Errorf("sms request body invalid")
	}
	phoneArray := strings.Split(phoneList, ",")
	fmt.Println(phoneArray)
	req.PhoneList = phoneArray
	req.ContentType = contentType
	req.Content = content
	req.SmsType = smsType
	if reserveTime != "" {
		req.ReserveTime = reserveTime
	}
	if scheduleCode != "" {
		req.ReserveTime = scheduleCode
	}
	image, err := c.FormFile("image")
	if err != nil && image != nil {
		return errorCommon.ErrorMsg(errorCommon.ErrBadParameter, errorCommon.Trace(), err.Error(), errorCommon.ErrFromClient)
	}
	if image != nil {
		file, err := image.Open()

		fileByte, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		imageEncode := base64.StdEncoding.EncodeToString(fileByte)
		req.Image = imageEncode
	}
	ctx := c.Request().Context()

	err = s.UseCase.Send(ctx, &req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, true)
}
