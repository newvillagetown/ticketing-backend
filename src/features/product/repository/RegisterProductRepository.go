package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"main/common/dbCommon/mysqlCommon"
	_interface "main/features/product/usecase/interface"
)

func NewRegisterProductRepository(gormDB *gorm.DB, tokenCollection *mongo.Collection) _interface.IRegisterProductRepository {
	return &RegisterProductRepository{GormDB: gormDB, TokenCollection: tokenCollection}
}

func (r *RegisterProductRepository) CreateProduct(ctx context.Context, productDTO mysqlCommon.GormProduct) error {
	result := mysqlCommon.GormDB.WithContext(ctx).Create(&productDTO)
	if result.RowsAffected == 0 || result.Error != nil {
		fmt.Println(result.RowsAffected)
		fmt.Println(result.Error)
		return fmt.Errorf("no row data")
	}
	return nil
}
