package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"main/common/dbCommon/mysqlCommon"
	_interface "main/features/product/usecase/interface"
	"time"
)

func NewGetProductRepository(tokenCollection *mongo.Collection) _interface.IGetProductRepository {
	return &GetProductRepository{TokenCollection: tokenCollection}
}

func (g *GetProductRepository) FindOneProduct(productID string) (mysqlCommon.GormProduct, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	var productDTO mysqlCommon.GormProduct
	result := mysqlCommon.GormDB.WithContext(ctx).Where("id = ?", productID).Find(&productDTO)
	if result.RowsAffected == 0 || result.Error != nil {
		return mysqlCommon.GormProduct{}, result.Error
	}
	return productDTO, nil
}
