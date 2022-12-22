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

func NewDeleteProductRepository(gormDB *gorm.DB, tokenCollection *mongo.Collection) _interface.IDeleteProductRepository {
	return &DeleteProductRepository{GormDB: gormDB, TokenCollection: tokenCollection}
}

func (d *DeleteProductRepository) FindOneAndDeleteUpdateProduct(ctx context.Context, productID string) error {
	result := d.GormDB.WithContext(ctx).Model(&mysqlCommon.GormProduct{}).Where("id = ?", productID).Update("is_deleted", true)

	if result.RowsAffected == 0 {
		return errorCommon.ErrorMsg(errorCommon.ErrBadParameter, errorCommon.Trace(), domain.ErrBadParamInput, errorCommon.ErrFromClient)
	}
	if result.Error != nil {
		return errorCommon.ErrorMsg(errorCommon.ErrInternalDB, errorCommon.Trace(), domain.ErrBadParamInput, errorCommon.ErrFromMysqlDB)
	}
	return nil
}
