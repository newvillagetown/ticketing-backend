package loggingCommon

import (
	"github.com/stretchr/testify/assert"
	"main/common/errorCommon"
	"testing"
	"time"
)

func TestMakeLog(t *testing.T) {
	t.Run("access log equal", func(t *testing.T) {
		now := time.Now()
		want := Log{
			Project:   "",
			Type:      "access",
			UserID:    "5623acad12",
			Url:       "/v0.1/test",
			Method:    "POST",
			Latency:   time.Since(now).Milliseconds(),
			HttpCode:  200,
			RequestID: "623acsdsda",
		}
		got := Log{}
		got.MakeLog("5623acad12", "/v0.1/test", "POST", now, 200, "623acsdsda")
		assert.Equal(t, got, want)
	})
}

func TestMakeErrorLog(t *testing.T) {
	t.Run("error log equal", func(t *testing.T) {

		requestParam := make(map[string]interface{}, 0)
		requestParam["test"] = "test01"
		resError := errorCommon.Err{
			HttpCode: 401,
			ErrType:  "AUTH_FAILED",
			Msg:      "auth failed",
			Trace:    "dev/src/test.go",
			From:     "client",
		}
		errInfo := ErrorInfo{
			Stack:     resError.Trace,
			ErrorType: resError.ErrType,
			From:      resError.From,
			Msg:       resError.Msg,
		}
		want := Log{
			Type:      "error",
			ErrorInfo: errInfo,
		}

		got := Log{}
		got.MakeErrorLog(resError)
		assert.Equal(t, got, want)

	})
}
