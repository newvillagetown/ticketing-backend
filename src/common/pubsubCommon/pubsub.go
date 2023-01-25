package pubsubCommon

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"golang.org/x/sync/errgroup"
	"main/common/dbCommon/mongodbCommon"
	"main/common/envCommon"
	"main/common/nCloudSmsCommon"
	"main/common/noticeCommon/productNotice"
	"time"
)

func InitPubSub() error {
	PubSubCh = gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(false, false),
	)
	// 받을 채널을 만든거다
	mongodbLogCh, err := PubSubCh.Subscribe(context.Background(), string(SubMongoDBLog))
	if err != nil {
		fmt.Println(err.Error())
	}
	productNoticeCh, err := PubSubCh.Subscribe(context.Background(), string(SubProductNotice))
	if err != nil {
		fmt.Println(err.Error())
	}
	//kakaoNotificationCh, err := PubSubCh.Subscribe(context.Background(), string(SubKakaoNotificationTalk))
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	naverSmsCh, err := PubSubCh.Subscribe(context.Background(), string(SubNaverSms))
	if err != nil {
		fmt.Println(err.Error())
	}

	go NaverSmsService(naverSmsCh)
	go LogProcess(mongodbLogCh)
	go GoogleRegisterNoticeProcess(productNoticeCh)

	return nil
}
func NaverSmsService(messages <-chan *message.Message) {
	for msg := range messages {
		data := &NaverSms{}
		err := json.Unmarshal(msg.Payload, data)
		if err != nil {
			fmt.Println(err.Error())
		}
		err = data.send()
		if err != nil {
			fmt.Println(err.Error())
		}

	}
}

// 이도영
func (n *NaverSms) send() error {
	g := new(errgroup.Group)
	now := time.Now()
	for i := 0; i < len(n.PhoneList); i++ {
		i := i
		g.Go(func() error {
			res, err := nCloudSmsCommon.NSmsSend(n.PhoneList[i], n.SmsType, n.ContentType, n.Content)
			fmt.Println(res)
			// DB 데이터 생성
			msgEvent := mongodbCommon.MessageEvent{
				Type:     string(SubNaverSms),
				State:    2,
				Occurred: envCommon.TimeToEpochMillis(now),
				NaverSms: mongodbCommon.NaverSms{
					Phone:       n.PhoneList[i],
					SmsType:     n.SmsType,
					Title:       n.Title,
					ContentType: n.ContentType,
					Content:     n.Content,
				},
				ResInfo: res,
			}
			if err != nil {
				//에러 처리
				fmt.Println("에러 처리 ㄱㄱ")
			}
			// DB 등록
			mongodbCommon.EventCollection.InsertOne(context.TODO(), msgEvent)

			return err
		})
	}
	if err := g.Wait(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// 구글 챗 노티 함수
func GoogleRegisterNoticeProcess(messages <-chan *message.Message) {
	for msg := range messages {
		data := &productNotice.ProductRegisterNotice{}
		now := time.Now()
		event := &mongodbCommon.Event{
			State:      true,
			OccurredAt: now.String(),
			Type:       string(ProductRegisterEventType),
		}
		err := json.Unmarshal(msg.Payload, data)
		if err != nil {
			event.ErrorMsg = err.Error()
			event.State = false
		}
		err = data.Send()
		if err != nil {
			event.ErrorMsg = err.Error()
			event.State = false
		}

		//이벤트 등록
		ctx := context.TODO()
		_, err = mongodbCommon.EventCollection.InsertOne(ctx, event)
		if err != nil {
			fmt.Println(err)
		}
		msg.Ack()
	}
}

// 메시지 처리 함수
func LogProcess(messages <-chan *message.Message) {
	for msg := range messages {
		data := &mongodbCommon.Log{}
		err := json.Unmarshal(msg.Payload, data)
		if err != nil {
			fmt.Println(err)
		}
		ctx := context.TODO()
		_, err = mongodbCommon.LogCollection.InsertOne(ctx, data)
		if err != nil {
			fmt.Println(err)
		}
		msg.Ack()
	}
}

// 메시지 발송 함수
// topic , msg 문자열로 받는다?
func PublishMessages(topic ISubscribeTopicType, msg interface{}, publisher message.Publisher) {
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
	}
	payload := message.NewMessage(watermill.NewUUID(), jsonMsg)
	if err = publisher.Publish(string(topic), payload); err != nil {
		fmt.Println(err)
	}
}
