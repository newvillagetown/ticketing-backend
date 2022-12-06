package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"main/common/dbCommon/mysqlCommon"
	_interface "main/features/product/usecase/interface"
)

func NewGetProductRepository(Conn *gorm.DB, tokenCollection *mongo.Collection) _interface.IGetProductRepository {
	return &GetProductRepository{Conn: Conn, TokenCollection: tokenCollection}
}

func (g *GetProductRepository) FindOneProduct(ctx context.Context, productID string) (mysqlCommon.GormProduct, error) {

	var productDTO mysqlCommon.GormProduct
	result := g.Conn.WithContext(ctx).Where("id = ?", productID).Find(&productDTO)
	if result.RowsAffected == 0 || result.Error != nil {
		return mysqlCommon.GormProduct{}, result.Error
	}
	return productDTO, nil
}
