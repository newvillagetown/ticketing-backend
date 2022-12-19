package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"main/common/dbCommon/mysqlCommon"
	_interface "main/features/product/usecase/interface"
)

func NewUpdateProductRepository(gormDB *gorm.DB, tokenCollection *mongo.Collection) _interface.IUpdateProductRepository {
	return &UpdateProductRepository{GormDB: gormDB, TokenCollection: tokenCollection}
}

func (u *UpdateProductRepository) FindOneProduct(ctx context.Context, productID string) (mysqlCommon.GormProduct, error) {
	var productDTO mysqlCommon.GormProduct
	result := mysqlCommon.GormDB.WithContext(ctx).Where("id = ?", productID).Find(&productDTO)
	if result.RowsAffected == 0 || result.Error != nil {
		return mysqlCommon.GormProduct{}, nil
	}
	return productDTO, nil
}

func (u *UpdateProductRepository) FindOneAndUpdateProduct(ctx context.Context, updatedProductDTO mysqlCommon.GormProduct) error {
	result := mysqlCommon.GormDB.WithContext(ctx).Save(&updatedProductDTO)
	if result.RowsAffected == 0 || result.Error != nil {
		return fmt.Errorf("no row data")
	}
	return nil
}
