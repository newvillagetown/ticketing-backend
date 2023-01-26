package _interface

import (
	"context"
	"main/features/message/domain/request"
)

type ISendNaverSmsUseCase interface {
	Send(ctx context.Context, req *request.ReqSendNaverSms) error
}
