package _interface

import (
	"context"
	"main/common/dbCommon/mongodbCommon"
)

type ISendNaverSmsRepository interface {
	SaveNaverSmsEvent(ctx context.Context, msgEvent mongodbCommon.NaverSmsEventDTO) error
}
