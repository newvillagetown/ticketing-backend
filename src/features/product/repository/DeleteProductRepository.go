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
	result, err := mysqlCommon.MysqlDB.Exec("update product set isDeleted = ? where id = ?", true, productID)
	if err != nil {
		return err
	}
	fmt.Println(result.RowsAffected())
	return nil
}
