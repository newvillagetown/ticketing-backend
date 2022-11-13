package logging

type AccessLog struct {
	Project   string `json:"project"`
	Type      string `json:"type"`
	UserID    string `json:"userID"`
	Url       string `json:"url"`
	Method    string `json:"method"`
	Latency   int    `json:"latency"`
	HttpCode  int    `json:"httpCode"`
	RequestID string `json:"requestID"`
}

type ErrorLog struct {
	Project      string                 `json:"project"`
	Type         string                 `json:"type"`
	UserID       string                 `json:"userID"`
	Url          string                 `json:"url"`
	Method       string                 `json:"method"`
	Latency      int                    `json:"latency"`
	HttpCode     int                    `json:"httpCode"`
	RequestID    string                 `json:"requestID"`
	Stack        string                 `json:"stack"`
	ErrorType    string                 `json:"errorType"`
	Err          string                 `json:"err"`
	From         string                 `json:"from"`
	RequestParam map[string]interface{} `json:"requestParam"`
}
