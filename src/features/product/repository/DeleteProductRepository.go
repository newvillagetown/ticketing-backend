package repository

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"main/common/dbCommon/mysqlCommon"
	_interface "main/features/product/usecase/interface"
)

func NewDeleteProductRepository(tokenCollection *mongo.Collection) _interface.IDeleteProductRepository {
	return &DeleteProductRepository{TokenCollection: tokenCollection}
}

func (d *DeleteProductRepository) FindOneAndDeleteUpdateProduct(productID string) error {
	fmt.Println(productID)
	result := mysqlCommon.GormDB.Model(&mysqlCommon.GormProduct{}).Where("id = ?", productID).Update("is_deleted", true)
	if result.RowsAffected == 0 || result.Error != nil {
		return result.Error
	}
	return nil
}
