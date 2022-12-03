package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"main/common/dbCommon/mysqlCommon"
	_interface "main/features/product/usecase/interface"
)

func NewGetProductRepository(tokenCollection *mongo.Collection) _interface.IGetProductRepository {
	return &GetProductRepository{TokenCollection: tokenCollection}
}

func (g *GetProductRepository) FindOneProduct(productID string) (mysqlCommon.GormProduct, error) {
	var productDTO mysqlCommon.GormProduct

	result := mysqlCommon.GormDB.Where("id = ?", productID).Find(&productDTO)
	if result.RowsAffected == 0 || result.Error != nil {
		return mysqlCommon.GormProduct{}, nil
	}
	return productDTO, nil
}
