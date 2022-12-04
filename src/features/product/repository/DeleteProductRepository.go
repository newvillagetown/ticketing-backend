package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"main/common/dbCommon/mysqlCommon"
	_interface "main/features/product/usecase/interface"
	"time"
)

func NewDeleteProductRepository(tokenCollection *mongo.Collection) _interface.IDeleteProductRepository {
	return &DeleteProductRepository{TokenCollection: tokenCollection}
}

func (d *DeleteProductRepository) FindOneAndDeleteUpdateProduct(productID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	result := mysqlCommon.GormDB.WithContext(ctx).Model(&mysqlCommon.GormProduct{}).Where("id = ?", productID).Update("is_deleted", true)
	if result.RowsAffected == 0 || result.Error != nil {
		return result.Error
	}
	return nil
}
