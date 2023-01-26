package usecase

import (
	"context"
	"golang.org/x/sync/errgroup"
	"main/common/dbCommon/mongodbCommon"
	"main/common/envCommon"
	"main/common/nCloudSmsCommon"
	"main/common/noticeCommon/naverSms"
	"main/common/pubsubCommon"
	"main/features/message/domain/request"
	_interface "main/features/message/usecase/interface"
	"time"
)

type SendNaverSmsUseCase struct {
	Repository     _interface.ISendNaverSmsRepository
	ContextTimeout time.Duration
}

func NewSendNaverSmsUseCase(repo _interface.ISendNaverSmsRepository, timeout time.Duration) _interface.ISendNaverSmsUseCase {
	return &SendNaverSmsUseCase{
		Repository:     repo,
		ContextTimeout: timeout,
	}
}

func (s *SendNaverSmsUseCase) Send(c context.Context, req *request.ReqSendNaverSms) error {

	g := new(errgroup.Group)
	now := time.Now()
	failedList := []string{}
	for i := 0; i < len(req.PhoneList); i++ {
		i := i
		g.Go(func() error {
			ctx, cancel := context.WithTimeout(c, s.ContextTimeout)
			defer cancel()
			// TODO 1. 네이버 메시지 전송
			res, err := nCloudSmsCommon.NSmsSend(req.PhoneList[i], req.SmsType, req.ContentType, req.Content)
			// DB 데이터 생성
			msgEvent := mongodbCommon.NaverSmsEventDTO{
				Type:     string(pubsubCommon.SubNaverSms),
				State:    2,
				Occurred: envCommon.TimeToEpochMillis(now),
				NaverSms: mongodbCommon.NaverSms{
					Phone:       req.PhoneList[i],
					SmsType:     req.SmsType,
					ContentType: req.ContentType,
					Content:     req.Content,
				},
			}
			if err != nil {
				msgEvent.Error = err.Error()
				failedList = append(failedList, req.PhoneList[i])
				return err
			}
			msgEvent.ResInfo = mongodbCommon.NaverSmsResDTO{
				StatusName:  res["statusName"].(string),
				RequestId:   res["requestId"].(string),
				RequestTime: res["requestTime"].(string),
				StatusCode:  res["statusCode"].(string),
			}
			// TODO 2. 이벤트 디비 히스토리 저장
			err = s.Repository.SaveNaverSmsEvent(ctx, msgEvent)
			if err != nil {
				return err
			}
			return err
		})
	}
	if err := g.Wait(); err != nil {
		// TODO 3. 에러에 대해서 구글 챗 메시지 전송
		googleChat := naverSms.NaverSmsFailedNotice{
			Content:   req.Content,
			PhoneList: failedList,
		}
		go googleChat.Send()
		return err
	}

	return nil
}
