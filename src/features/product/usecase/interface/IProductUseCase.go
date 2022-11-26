package _interface

import "main/features/product/model/request"

type IRegisterProductUseCase interface {
	Register(req request.ReqRegisterProduct) error
}
