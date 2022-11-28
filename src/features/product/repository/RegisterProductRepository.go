package repository

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"main/common/dbCommon/mysqlCommon"
	_interface "main/features/product/usecase/interface"
)

func NewRegisterProductRepository(tokenCollection *mongo.Collection) _interface.IRegisterProductRepository {
	return &RegisterProductRepository{TokenCollection: tokenCollection}
}

func (r *RegisterProductRepository) CreateProduct(productDTO mysqlCommon.Product) error {
	result, err := mysqlCommon.MysqlDB.Exec("INSERT INTO product(id, created, lastUpdated, isDeleted, name, description, category, perAmount, totalCount, restCount,startDate, endDate,imgUrl) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)",
		productDTO.ID, productDTO.Created, productDTO.LastUpdated, productDTO.IsDeleted, productDTO.Name, productDTO.Description, productDTO.Category, productDTO.PerAmount, productDTO.TotalCount, productDTO.RestCount, productDTO.StartDate, productDTO.EndDate, productDTO.ImgUrl)
	if err != nil {
		fmt.Println(err)
		return err
	}
	n, err := result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row inserted.")
	}
	return nil
}
