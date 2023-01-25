package pubsubCommon

import "github.com/ThreeDotsLabs/watermill/pubsub/gochannel"

type NaverSms struct {
	PhoneList   []string
	Title       string //subject	Optional	String	기본 메시지 제목	LMS, MMS에서만 사용 가능
	ContentType string //COMM 일반 메시지, AD 광고 메시지
	Content     string
}

const (
	SubMongoDBLog            = ISubscribeTopicType("mongodb.log")
	SubProductNotice         = ISubscribeTopicType("product.notice")
	SubKakaoNotificationTalk = ISubscribeTopicType("kakao.notification")
	SubKakaoFriendTalk       = ISubscribeTopicType("kakao.friend")
	SubNaverSms              = ISubscribeTopicType("naver.sms")
)

type EventType string

const (
	ProductRegisterEventType = EventType("PRODUCT_REGISTER")
)

var PubSubCh *gochannel.GoChannel

type ISubscribeTopicType string
