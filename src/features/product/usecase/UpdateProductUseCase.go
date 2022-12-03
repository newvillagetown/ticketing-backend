package usecase

import (
	"main/features/product/model/request"
	_interface "main/features/product/usecase/interface"
)

type UpdateProductUseCase struct {
	Repository _interface.IUpdateProductRepository
}

func NewUpdateProductUseCase(repo _interface.IUpdateProductRepository) _interface.IUpdateProductUseCase {
	return &UpdateProductUseCase{
		Repository: repo,
	}
}

func (u *UpdateProductUseCase) Update(req request.ReqUpdateProduct) error {
	productDTO, err := u.Repository.FindOneProduct(req.ProductID)
	if err != nil {
		return err
	}
	updatedProductDTO := ConvertUpdateProductNewProductDTO(req, productDTO)
	err = u.Repository.FindOneAndUpdateProduct(updatedProductDTO)
	if err != nil {
		return err
	}
	return nil
}
