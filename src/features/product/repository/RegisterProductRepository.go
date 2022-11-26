package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	_interface "main/features/product/usecase/interface"
)

func NewRegisterProductRepository(tokenCollection *mongo.Collection) _interface.IRegisterProductRepository {
	return &RegisterProductRepository{TokenCollection: tokenCollection}
}

func (r *RegisterProductRepository) CreateProduct() error {
	return nil
}
