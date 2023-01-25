package mongodbCommon

import (
	"time"
)

type RefreshToken struct {
	Created   time.Time `bson:"created"`
	Email     string    `bson:"email"`
	Token     string    `bson:"token"`
	IsDeleted bool      `bson:"isDeleted"`
}

type Log struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	Project   string    `json:"project"bson:"project"`
	Type      string    `json:"type" bson:"type"`
	UserID    string    `json:"userID" bson:"user_id"`
	Url       string    `json:"url" bson:"url"`
	Method    string    `json:"method" bson:"method"`
	Latency   int64     `json:"latency" bson:"latency"`
	HttpCode  int       `json:"httpCode" bson:"httpCode"`
	RequestID string    `json:"requestID" bson:"request_id"`
	ErrorInfo ErrorInfo `json:"errorInfo,omitempty" bson:"error_info"`
}

type ErrorInfo struct {
	Stack        string                 `json:"stack" bson:"stack"`
	ErrorType    string                 `json:"errorType" bson:"error_type"`
	Msg          string                 `json:"msg" bson:"msg"`
	From         string                 `json:"from" bson:"from"`
	RequestParam map[string]interface{} `json:"requestParam" bson:"request_param"`
}

type Event struct {
	ID         string `json:"id" bson:"_id,omitempty"`
	State      bool   `json:"state" bson:"state"`
	Type       string `json:"type" bson:"type"`
	OccurredAt string `json:"occurredAt" bson:"occurred_at"`
	ErrorMsg   string `json:"errorMsg" bson:"error_msg,omitempty"`
}

type MessageEvent struct {
	ID       string                 `json:"id" bson:"_id,omitempty"`
	Type     string                 `json:"type" bson:"type"`
	State    int                    `json:"state" bson:"state"`
	Occurred int64                  `json:"occurred" bson:"occurred"`
	NaverSms NaverSms               `json:"naverSms,omitempty" bson:"naverSms,omitempty"`
	ReqInfo  map[string]interface{} `json:"reqInfo" bson:"reqInfo"`
	ResInfo  map[string]interface{} `json:"resInfo" bson:"resInfo"`
	Error    string                 `json:"error,omitempty" bson:"error,omitempty"`
}

type NaverSms struct {
	Phone           string
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
