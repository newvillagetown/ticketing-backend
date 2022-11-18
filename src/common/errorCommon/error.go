package errorCommon

import (
	"fmt"
	"strings"
)

// 에러 메시지 생성할 때 사용
func NewError(errType ErrType, trace string, msg string, from IErrFrom) error {
	return fmt.Errorf("%s|%s|%s|%s", errType, trace, msg, from)
}

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
