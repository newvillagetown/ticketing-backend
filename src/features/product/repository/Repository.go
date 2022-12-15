package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type RegisterProductRepository struct {
	GormDB          *gorm.DB
	TokenCollection *mongo.Collection
}

type GetProductRepository struct {
	GormDB          *gorm.DB
	TokenCollection *mongo.Collection
}

type GetsProductRepository struct {
	GormDB          *gorm.DB
	TokenCollection *mongo.Collection
}

type DeleteProductRepository struct {
	GormDB          *gorm.DB
	TokenCollection *mongo.Collection
}

type UpdateProductRepository struct {
	GormDB          *gorm.DB
	TokenCollection *mongo.Collection
}
