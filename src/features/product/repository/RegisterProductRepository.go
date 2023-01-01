package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"main/common/dbCommon/mysqlCommon"
	"main/common/errorCommon"
	"main/features/product/domain"
	_interface "main/features/product/usecase/interface"
)

func NewRegisterProductRepository(gormDB *gorm.DB, tokenCollection *mongo.Collection) _interface.IRegisterProductRepository {
	return &RegisterProductRepository{GormDB: gormDB, TokenCollection: tokenCollection}
}

func (r *RegisterProductRepository) CreateProduct(ctx context.Context, productDTO mysqlCommon.GormProduct) error {
	result := r.GormDB.WithContext(ctx).Create(&productDTO)
	if result.RowsAffected == 0 {
		return errorCommon.ErrorMsg(errorCommon.ErrBadParameter, errorCommon.Trace(), domain.ErrBadParamInput, errorCommon.ErrFromClient)
	}
	if result.Error != nil {
		return errorCommon.ErrorMsg(errorCommon.ErrInternalDB, errorCommon.Trace(), result.Error.Error(), errorCommon.ErrFromMysqlDB)
	}
	return nil
}
