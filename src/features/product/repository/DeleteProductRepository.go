package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"main/common/dbCommon/mysqlCommon"
	_interface "main/features/product/usecase/interface"
)

func NewDeleteProductRepository(gormDB *gorm.DB, tokenCollection *mongo.Collection) _interface.IDeleteProductRepository {
	return &DeleteProductRepository{GormDB: gormDB, TokenCollection: tokenCollection}
}

func (d *DeleteProductRepository) FindOneAndDeleteUpdateProduct(ctx context.Context, productID string) error {
	result := d.GormDB.WithContext(ctx).Model(&mysqlCommon.GormProduct{}).Where("id = ?", productID).Update("is_deleted", true)
	if result.RowsAffected == 0 || result.Error != nil {
		return fmt.Errorf("no row data")
	}
	return nil
}
