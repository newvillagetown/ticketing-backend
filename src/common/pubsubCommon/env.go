package pubsubCommon

import "github.com/ThreeDotsLabs/watermill/pubsub/gochannel"

/*
MMS에서만 사용 가능
- 공백 사용 불가
- *.jpg, *.jpeg 확장자를 가진 파일 이름
- 최대 40자
*.jpg, *.jpeg 이미지를 Base64로 인코딩한 값
- 원 파일 기준 최대 300Kbyte
- 파일 명 최대 40자
- 해상도 최대 1500 * 1440
*/

type NaverSms struct {
	PhoneList       []string
	SmsType         string // sms/lms/mms
	Title           string // subject	Optional	String	기본 메시지 제목	LMS, MMS에서만 사용 가능
	ContentType     string // COMM 일반 메시지, AD 광고 메시지
	Content         string // Content
	File            File   //mms에서만 가능
	ReserveTime     string //메시지 발송 예약 일시 (yyyy-MM-dd HH:mm)
	ReserveTimeZone string // 예약 일시
	ScheduleCode    string // 문자 메시지 주기
}
type File struct {
	Name string //.jpg, jpeg 확장자 가능, 최대 40자, 공백 x
	Body string // 최대 300kbyte, base64인코딩한 값
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
