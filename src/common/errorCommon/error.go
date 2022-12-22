package errorCommon

import (
	"fmt"
	"runtime"
	"strings"
)

func ErrorParsing(data string) Err {
	slice := strings.Split(data, "|")
	result := Err{
		HttpCode: ErrHttpCode[slice[0]],
		ErrType:  slice[0],
		Trace:    slice[1],
		Msg:      slice[2],
		From:     slice[3],
	}
	return result
}

func ErrorMsg(errType ErrType, trace string, msg string, from IErrFrom) error {
	return fmt.Errorf("%s|%s|%s|%s", errType, trace, msg, from)
}

func (e ErrType) New(errType string, msg string) *ResError {
	return &ResError{ErrType: errType, Msg: msg}
}

func Trace() string {
	pc, _, _, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	_, line := runtime.FuncForPC(pc).FileLine(pc)
	return fmt.Sprintf("%s.L%d", funcName, line)
}
