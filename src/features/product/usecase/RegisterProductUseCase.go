package usecase

import _interface "main/features/product/usecase/interface"

type RegisterProductUseCase struct {
	Repository _interface.IRegisterProductRepository
}

func NewRegisterProductUseCase(repo _interface.IRegisterProductRepository) _interface.IRegisterProductUseCase {
	return &RegisterProductUseCase{
		Repository: repo,
	}
}

func (r *RegisterProductUseCase) Register() error {

	return nil
}
