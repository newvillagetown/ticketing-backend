package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"main/common/dbCommon/mysqlCommon"
	_interface "main/features/product/usecase/interface"
	"time"
)

func NewUpdateProductRepository(tokenCollection *mongo.Collection) _interface.IUpdateProductRepository {
	return &UpdateProductRepository{TokenCollection: tokenCollection}
}

func (u *UpdateProductRepository) FindOneProduct(productID string) (mysqlCommon.GormProduct, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	var productDTO mysqlCommon.GormProduct
	result := mysqlCommon.GormDB.WithContext(ctx).Where("id = ?", productID).Find(&productDTO)
	if result.RowsAffected == 0 || result.Error != nil {
		return mysqlCommon.GormProduct{}, nil
	}
	return productDTO, nil
}

func (u *UpdateProductRepository) FindOneAndUpdateProduct(updatedProductDTO mysqlCommon.GormProduct) error {
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	result := mysqlCommon.GormDB.WithContext(ctx).Save(&updatedProductDTO)
	if result.RowsAffected == 0 || result.Error != nil {
		return result.Error
	}
	return nil
}
