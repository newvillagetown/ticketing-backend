package loggingCommon

import (
	"github.com/rs/zerolog/log"
	"main/common/envCommon"
	"main/common/errorCommon"
	"time"
)

type Logging interface {
	MakeLog() error
	MakeErrorLog(requestParam map[string]interface{}) error
}

type Log struct {
	Project   string    `json:"project"`
	Type      string    `json:"type"`
	UserID    string    `json:"userID"`
	Url       string    `json:"url"`
	Method    string    `json:"method"`
	Latency   int64     `json:"latency"`
	HttpCode  int       `json:"httpCode"`
	RequestID string    `json:"requestID"`
	ErrorInfo ErrorInfo `json:"errorInfo,omitempty"`
}

type ErrorInfo struct {
	Stack        string                 `json:"stack"`
	ErrorType    string                 `json:"errorType"`
	Msg          string                 `json:"msg"`
	From         string                 `json:"from"`
	RequestParam map[string]interface{} `json:"requestParam"`
}

func (l *Log) MakeLog(userID string, url string, method string, startTime time.Time, httpCode int, requestID string) error {
	l.Project = envCommon.Env.Project
	l.Type = "access"
	l.UserID = userID
	l.Url = url
	l.Method = method
	l.Latency = time.Since(startTime).Milliseconds()
	l.HttpCode = httpCode
	l.RequestID = requestID

	return nil
}
func (l *Log) MakeErrorLog(requestParam map[string]interface{}, resError errorCommon.Err) error {
	l.Type = "error"
	errInfo := ErrorInfo{
		RequestParam: requestParam,
		Stack:        resError.Trace,
		ErrorType:    resError.ErrType,
		From:         resError.From,
		Msg:          resError.Msg,
	}
	l.ErrorInfo = errInfo
	return nil
}

func (l *Log) InfoLog() error {
	log.Info().Str("project", l.Project).
		Str("type", l.Type).
		Str("userID", l.UserID).
		Str("url", l.Url).
		Str("method", l.Method).
		Int64("latency", l.Latency).
		Int("httpCode", l.HttpCode).
		Str("requestID", l.RequestID)
	return nil
}
func (l *Log) ErrorLog() error {
	log.Error().Str("project", l.Project).
		Str("type", l.Type).
		Str("userID", l.UserID).
		Str("url", l.Url).
		Str("method", l.Method).
		Int64("latency", l.Latency).
		Int("httpCode", l.HttpCode).
		Str("requestID", l.RequestID)
	return nil
}
