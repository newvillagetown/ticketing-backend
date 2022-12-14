package pubsubCommon

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"main/common/dbCommon/mongodbCommon"
	"main/common/noticeCommon/productNotice"
	"time"
)

const (
	SubMongoDBLog    = ISubscribeTopicType("mongodb.log")
	SubProductNotice = ISubscribeTopicType("product.notice")
)

type EventType string

const (
	ProductRegisterEventType = EventType("PRODUCT_REGISTER")
)

var PubSubCh *gochannel.GoChannel

type ISubscribeTopicType string

func InitPubSub() error {
	PubSubCh = gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(false, false),
	)
	// 받을 채널을 만든거다
	mongodbLogCh, err := PubSubCh.Subscribe(context.Background(), string(SubMongoDBLog))
	if err != nil {
		panic(err)
	}
	productNoticeCh, err := PubSubCh.Subscribe(context.Background(), string(SubProductNotice))
	go LogProcess(mongodbLogCh)
	go GoogleRegisterNoticeProcess(productNoticeCh)

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
