package mongodbCommon

import "time"

type RefreshToken struct {
	Created   time.Time `bson:"created"`
	Email     string    `bson:"email"`
	Token     string    `bson:"token"`
	IsDeleted bool      `bson:"isDeleted"`
}

type Log struct {
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
