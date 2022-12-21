package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"main/common/dbCommon/mysqlCommon"
	"main/common/errorCommon"
	_interface "main/features/product/usecase/interface"
)

func NewGetProductRepository(gormDB *gorm.DB, tokenCollection *mongo.Collection) _interface.IGetProductRepository {
	return &GetProductRepository{GormDB: gormDB, TokenCollection: tokenCollection}
}

func (g *GetProductRepository) FindOneProduct(ctx context.Context, productID string) (mysqlCommon.GormProduct, error) {

	var productDTO mysqlCommon.GormProduct
	result := g.GormDB.WithContext(ctx).Where("id = ?", productID).Find(&productDTO)
	if result.RowsAffected == 0 {
		return mysqlCommon.GormProduct{}, errorCommon.ErrBadParamInput
	}
	if result.Error != nil {
		return mysqlCommon.GormProduct{}, result.Error
	}
	return productDTO, nil
}
