package logging

import (
	"github.com/rs/zerolog/log"
	"main/common/env"
	"time"
)

type Logging interface {
	MakeLog() error
	MakeErrorLog() error
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
	Err          string                 `json:"err"`
	From         string                 `json:"from"`
	RequestParam map[string]interface{} `json:"requestParam"`
}

func (l *Log) MakeLog(userID string, url string, method string, startTime time.Time, httpCode int, requestID string) error {
	l.Project = env.Env.Project
	l.Type = "access"
	l.UserID = userID
	l.Url = url
	l.Method = method
	l.Latency = time.Since(startTime).Milliseconds()
	l.HttpCode = httpCode
	l.RequestID = requestID
	return nil
}
func (l *Log) MakeErrorLog() error {
	l.Type = "error"

	return nil
}

func (l *Log) InfoLog() error {
	log.Info().Str("project", l.Project)
	//	log.Info().Str("URI", url).Int("state", 200).Int64("latency", time.Since(startTime).Milliseconds()).Str("METHOD", "POST").Msg("hi")
	return nil
}
func (l *Log) ErrorLog() error {
	log.Error().Str("project", l.Project)
	return nil
}
