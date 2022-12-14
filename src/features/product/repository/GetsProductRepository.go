package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"main/common/dbCommon/mysqlCommon"
	"main/common/errorCommon"
	_interface "main/features/product/usecase/interface"
)

func NewGetsProductRepository(gormDB *gorm.DB, tokenCollection *mongo.Collection) _interface.IGetsProductRepository {
	return &GetsProductRepository{GormDB: gormDB, TokenCollection: tokenCollection}
}

func (g *GetsProductRepository) FindProduct(ctx context.Context) ([]mysqlCommon.GormProduct, error) {
	var productsDTO []mysqlCommon.GormProduct
	result := g.GormDB.WithContext(ctx).Where("is_deleted", false).Find(&productsDTO)
	if result.Error != nil {
		return nil, errorCommon.ErrorMsg(errorCommon.ErrInternalDB, errorCommon.Trace(), result.Error.Error(), errorCommon.ErrFromMysqlDB)
	}

	return productsDTO, nil
}
