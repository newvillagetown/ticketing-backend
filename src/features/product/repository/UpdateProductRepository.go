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

func NewUpdateProductRepository(gormDB *gorm.DB, tokenCollection *mongo.Collection) _interface.IUpdateProductRepository {
	return &UpdateProductRepository{GormDB: gormDB, TokenCollection: tokenCollection}
}

func (u *UpdateProductRepository) FindOneProduct(ctx context.Context, productID string) (mysqlCommon.GormProduct, error) {
	var productDTO mysqlCommon.GormProduct
	result := u.GormDB.WithContext(ctx).Where("id = ?", productID).Find(&productDTO)
	if result.RowsAffected == 0 {
		return mysqlCommon.GormProduct{}, errorCommon.ErrorMsg(errorCommon.ErrBadParameter, errorCommon.Trace(), domain.ErrBadParamInput, errorCommon.ErrFromClient)
	}
	if result.Error != nil {
		return mysqlCommon.GormProduct{}, errorCommon.ErrorMsg(errorCommon.ErrInternalDB, errorCommon.Trace(), result.Error.Error(), errorCommon.ErrFromMysqlDB)
	}
	return productDTO, nil
}

func (u *UpdateProductRepository) FindOneAndUpdateProduct(ctx context.Context, updatedProductDTO mysqlCommon.GormProduct) error {
	result := u.GormDB.WithContext(ctx).Save(&updatedProductDTO)
	if result.RowsAffected == 0 {
		return errorCommon.ErrorMsg(errorCommon.ErrBadParameter, errorCommon.Trace(), domain.ErrBadParamInput, errorCommon.ErrFromClient)
	}
	if result.Error != nil {
		return errorCommon.ErrorMsg(errorCommon.ErrInternalDB, errorCommon.Trace(), result.Error.Error(), errorCommon.ErrFromMysqlDB)
	}
	return nil
}
