package handler

import (
	"github.com/labstack/echo/v4"
	"main/common/dbCommon/mongodbCommon"
	"main/common/dbCommon/mysqlCommon"
	"main/features/product/repository"
	"main/features/product/usecase"
)

func NewProductHandler(c *echo.Echo) {
	NewRegisterProductHandler(c, usecase.NewRegisterProductUseCase(repository.NewRegisterProductRepository(mysqlCommon.GormDB, mongodbCommon.TokenCollection), mysqlCommon.DBTimeOut))
	NewDeleteProductHandler(c, usecase.NewDeleteProductUseCase(repository.NewDeleteProductRepository(mysqlCommon.GormDB, mongodbCommon.TokenCollection), mysqlCommon.DBTimeOut))
	NewGetsProductHandler(c, usecase.NewGetsProductUseCase(repository.NewGetsProductRepository(mysqlCommon.GormDB, mongodbCommon.TokenCollection), mysqlCommon.DBTimeOut))
	NewGetProductHandler(c, usecase.NewGetProductUseCase(repository.NewGetProductRepository(mysqlCommon.GormDB, mongodbCommon.TokenCollection), mysqlCommon.DBTimeOut))
	NewUpdateProductHandler(c, usecase.NewUpdateProductUseCase(repository.NewUpdateProductRepository(mysqlCommon.GormDB, mongodbCommon.TokenCollection), mysqlCommon.DBTimeOut))
}
