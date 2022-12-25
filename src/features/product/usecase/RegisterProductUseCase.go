package usecase

import (
	"context"
	"main/features/product/domain/request"
	_interface "main/features/product/usecase/interface"
	"time"
)

type RegisterProductUseCase struct {
	Repository     _interface.IRegisterProductRepository
	ContextTimeout time.Duration
}

func NewRegisterProductUseCase(repo _interface.IRegisterProductRepository, timeout time.Duration) _interface.IRegisterProductUseCase {
	return &RegisterProductUseCase{
		Repository:     repo,
		ContextTimeout: timeout,
	}
}

func (r *RegisterProductUseCase) Register(c context.Context, req request.ReqRegisterProduct) error {
	ctx, cancel := context.WithTimeout(c, r.ContextTimeout)
	defer cancel()
	productDTO := ConvertToRegisterProductDTO(req)
	err := r.Repository.CreateProduct(ctx, productDTO)
	if err != nil {
		return err
	}
	//TODO 제품 등록시 구글 챗으로 전송
	googleSend := MakeProductRegisterNotice(productDTO)
	googleSend.Send()
	return nil
}
