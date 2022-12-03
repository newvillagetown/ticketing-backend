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

func (r *RegisterProductRepository) CreateProduct(productDTO mysqlCommon.GormProduct) error {
	result := mysqlCommon.GormDB.Create(&productDTO)
	if result.RowsAffected == 0 || result.Error != nil {
		fmt.Println(result.RowsAffected)
		fmt.Println(result.Error)
		return result.Error
	}
	return nil
}
