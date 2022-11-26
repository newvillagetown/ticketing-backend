package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	_interface "main/features/product/usecase/interface"
)

func NewGetsProductRepository(tokenCollection *mongo.Collection) _interface.IGetsProductRepository {
	return &GetsProductRepository{TokenCollection: tokenCollection}
}

func (g *GetsProductRepository) FindProduct() error {
	return nil
}
