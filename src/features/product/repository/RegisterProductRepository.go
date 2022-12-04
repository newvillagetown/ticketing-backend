package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"main/common/dbCommon/mysqlCommon"
	_interface "main/features/product/usecase/interface"
	"time"
)

func NewRegisterProductRepository(tokenCollection *mongo.Collection) _interface.IRegisterProductRepository {
	return &RegisterProductRepository{TokenCollection: tokenCollection}
}

func (r *RegisterProductRepository) CreateProduct(productDTO mysqlCommon.GormProduct) error {
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	result := mysqlCommon.GormDB.WithContext(ctx).Create(&productDTO)
	if result.RowsAffected == 0 || result.Error != nil {
		fmt.Println(result.RowsAffected)
		fmt.Println(result.Error)
		return result.Error
	}
	return nil
}
