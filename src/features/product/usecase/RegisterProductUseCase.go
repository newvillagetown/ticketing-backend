package usecase

import (
	"main/features/product/model/request"
	_interface "main/features/product/usecase/interface"
)

type RegisterProductUseCase struct {
	Repository _interface.IRegisterProductRepository
}

func NewRegisterProductUseCase(repo _interface.IRegisterProductRepository) _interface.IRegisterProductUseCase {
	return &RegisterProductUseCase{
		Repository: repo,
	}
}

func (r *RegisterProductUseCase) Register(req request.ReqRegisterProduct) error {
	productDTO := ConvertToRegisterProductDTO(req)
	err := r.Repository.CreateProduct(productDTO)
	if err != nil {
		return err
	}
	return nil
}
