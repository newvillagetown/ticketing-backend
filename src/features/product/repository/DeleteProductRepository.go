package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	_interface "main/features/product/usecase/interface"
)

func NewDeleteProductRepository(tokenCollection *mongo.Collection) _interface.IDeleteProductRepository {
	return &DeleteProductRepository{TokenCollection: tokenCollection}
}

func (d *DeleteProductRepository) FindOneAndDeleteUpdateProduct() error {
	return nil
}
