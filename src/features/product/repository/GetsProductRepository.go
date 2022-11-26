package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"main/common/dbCommon/mysqlCommon"
	_interface "main/features/product/usecase/interface"
)

func NewGetsProductRepository(tokenCollection *mongo.Collection) _interface.IGetsProductRepository {
	return &GetsProductRepository{TokenCollection: tokenCollection}
}

func (g *GetsProductRepository) FindProduct() ([]mysqlCommon.Product, error) {

	rows, err := mysqlCommon.MysqlDB.Query("SELECT id,created,lastUpdated,isDeleted,name,description,category,perAmount,totalCount,restCount,startDate,endDate FROM product where isDeleted = false")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	productList := make([]mysqlCommon.Product, 0)
	for rows.Next() {
		productDTO := mysqlCommon.Product{}
		err := rows.Scan(
			&productDTO.ID, &productDTO.Created, &productDTO.LastUpdated, &productDTO.IsDeleted, &productDTO.Name, &productDTO.Description, &productDTO.Category, &productDTO.PerAmount,
			&productDTO.TotalCount, &productDTO.RestCount, &productDTO.StartDate, &productDTO.EndDate)

		if err != nil {
			return nil, err
		}
		productList = append(productList, productDTO)
	}

	return productList, nil
}
