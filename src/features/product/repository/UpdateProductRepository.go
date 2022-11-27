package repository

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"main/common/dbCommon/mysqlCommon"
	_interface "main/features/product/usecase/interface"
)

func NewUpdateProductRepository(tokenCollection *mongo.Collection) _interface.IUpdateProductRepository {
	return &UpdateProductRepository{TokenCollection: tokenCollection}
}

func (u *UpdateProductRepository) FindOneProduct(productID string) (mysqlCommon.Product, error) {
	productDTO := mysqlCommon.Product{}
	err := mysqlCommon.MysqlDB.QueryRow("SELECT id,created,lastUpdated,isDeleted,name,description,category,perAmount,totalCount,restCount,startDate,endDate FROM product WHERE id = ?", productID).Scan(
		&productDTO.ID, &productDTO.Created, &productDTO.LastUpdated, &productDTO.IsDeleted, &productDTO.Name, &productDTO.Description, &productDTO.Category, &productDTO.PerAmount,
		&productDTO.TotalCount, &productDTO.RestCount, &productDTO.StartDate, &productDTO.EndDate)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			return mysqlCommon.Product{}, err
		}
		return mysqlCommon.Product{}, nil
	}
	return productDTO, nil
}

func (u *UpdateProductRepository) FindOneAndUpdateProduct(productDTO mysqlCommon.Product) error {
	result, err := mysqlCommon.MysqlDB.Exec("update product set created = ?, lastUpdated = ?, isDeleted = ?, name = ?, description = ?, category = ?, perAmount = ?, totalCount = ?, restCount = ?, startDate = ?, endDate = ? where id = ?",
		productDTO.Created, productDTO.LastUpdated, productDTO.IsDeleted, productDTO.Name, productDTO.Description, productDTO.Category, productDTO.PerAmount, productDTO.TotalCount, productDTO.RestCount, productDTO.StartDate, productDTO.EndDate, productDTO.ID)
	if err != nil {
		return err
	}
	n, _ := result.RowsAffected()
	if n == 0 {
		return fmt.Errorf("sql: no rows in result set")
	}
	return nil
}
