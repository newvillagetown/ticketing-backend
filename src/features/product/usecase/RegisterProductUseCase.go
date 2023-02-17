package usecase

import (
	"context"
	"main/common/pubsubCommon"
	"main/features/product/domain/request"
	"main/features/product/usecase/interface"
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
	googleSend := MakeProductRegisterNotice(productDTO)
	pubsubCommon.PublishMessages(pubsubCommon.SubProductNotice, googleSend, pubsubCommon.PubSubCh)
	return nil
}
