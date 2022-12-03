package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"main/common/dbCommon/mysqlCommon"
	_interface "main/features/product/usecase/interface"
)

func NewGetsProductRepository(tokenCollection *mongo.Collection) _interface.IGetsProductRepository {
	return &GetsProductRepository{TokenCollection: tokenCollection}
}

func (g *GetsProductRepository) FindProduct() ([]mysqlCommon.GormProduct, error) {
	var productsDTO []mysqlCommon.GormProduct
	result := mysqlCommon.GormDB.Find(&productsDTO)
	if result.Error != nil {
		return nil, result.Error
	}
	return productsDTO, nil
}
