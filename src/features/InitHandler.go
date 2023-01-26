package features

import (
	"github.com/labstack/echo/v4"
	"main/common/pubsubCommon"
	messageHandler "main/features/message/handler"
	googleOAuthHandler "main/features/oauth/google/handler"
	productHandler "main/features/product/handler"
	userHandler "main/features/user/handler"
	"net/http"
)

func InitHandler(e *echo.Echo) error {
	//elb 헬스체크용
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	e.GET("/test", func(c echo.Context) error {
		msg := pubsubCommon.NaverSms{
			PhoneList:   []string{"01051105508", "01029121993"},
			ContentType: "COMM",
			Content:     "메시지 오니? -테스트",
			SmsType:     "SMS",
		}
		pubsubCommon.PublishMessages(pubsubCommon.SubNaverSms, msg, pubsubCommon.PubSubCh)
		return c.NoContent(http.StatusOK)
	})

	//인증 핸들러 초기화
	googleOAuthHandler.NewGoogleOAuthHandler(e)

	//기능 핸들러 초기화
	productHandler.NewProductHandler(e)
	userHandler.NewUserHandler(e)
	messageHandler.NewNaverSmsHandler(e)

	return nil
}
