package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	_interface "main/features/product/usecase/interface"
)

func NewGetProductRepository(tokenCollection *mongo.Collection) _interface.IGetProductRepository {
	return &GetProductRepository{TokenCollection: tokenCollection}
}

func (g *GetProductRepository) FindOneProduct() error {

	return nil
}
