package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	_interface "main/features/product/usecase/interface"
)

func NewUpdateProductRepository(tokenCollection *mongo.Collection) _interface.IUpdateProductRepository {
	return &UpdateProductRepository{TokenCollection: tokenCollection}
}

func (u *UpdateProductRepository) FindOneAndUpdateProduct() error {

	return nil
}
