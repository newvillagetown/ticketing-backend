package request

/*
PhoneList:   []string{"01051105508", "01029121993"},
ContentType: "COMM",
Content:     "메시지 오니? -테스트",
SmsType:     "SMS",
"reserveTime": "yyyy-MM-dd HH:mm",

	"reserveTimeZone": "string",
	"scheduleCode": "string"
*/
type ReqSendNaverSms struct {
	PhoneList    []string `json:"phoneList"`
	ContentType  string   `json:"contentType"`
	Content      string   `json:"content"`
	SmsType      string   `json:"smsType"`
	ReserveTime  string   `json:"reserveTime,omitempty"`
	ScheduleCode string   `json:"scheduleCode,omitempty"`
	Image        string   `json:"image,omitempty"`
}
