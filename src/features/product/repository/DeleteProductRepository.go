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
	result, err := mysqlCommon.MysqlDB.Exec("update product set isDeleted = ? where id = ? and isDeleted = ?", true, productID, false)
	if err != nil {
		return err
	}
	cnt, err := result.RowsAffected()
	if cnt == 0 {
		return fmt.Errorf("sql: no rows in result set")
	}
	return nil
}
