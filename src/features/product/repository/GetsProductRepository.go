package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"main/common/dbCommon/mysqlCommon"
	_interface "main/features/product/usecase/interface"
	"time"
)

func NewGetsProductRepository(tokenCollection *mongo.Collection) _interface.IGetsProductRepository {
	return &GetsProductRepository{TokenCollection: tokenCollection}
}

func (g *GetsProductRepository) FindProduct() ([]mysqlCommon.GormProduct, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	var productsDTO []mysqlCommon.GormProduct
	result := mysqlCommon.GormDB.WithContext(ctx).Where("is_deleted", false).Find(&productsDTO)
	if result.Error != nil {
		return nil, result.Error
	}
	return productsDTO, nil
}
