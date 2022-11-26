package repository

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"main/common/dbCommon/mysqlCommon"
	_interface "main/features/product/usecase/interface"
)

func NewGetProductRepository(tokenCollection *mongo.Collection) _interface.IGetProductRepository {
	return &GetProductRepository{TokenCollection: tokenCollection}
}

func (g *GetProductRepository) FindOneProduct(productID string) (mysqlCommon.Product, error) {
	productDTO := mysqlCommon.Product{}
	err := mysqlCommon.MysqlDB.QueryRow("SELECT id,created,lastUpdated,isDeleted,name,description,category,perAmount,totalCount,restCount,startDate,endDate FROM product WHERE id = ?", productID).Scan(
		&productDTO.ID, &productDTO.Created, &productDTO.LastUpdated, &productDTO.IsDeleted, &productDTO.Name, &productDTO.Description, &productDTO.Category, &productDTO.PerAmount,
		&productDTO.TotalCount, &productDTO.RestCount, &productDTO.StartDate, &productDTO.EndDate)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			fmt.Println(err)
			return mysqlCommon.Product{}, err
		}
		return mysqlCommon.Product{}, nil
	}
	return productDTO, nil
}
