package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type RegisterProductRepository struct {
	TokenCollection *mongo.Collection
}

type GetProductRepository struct {
	Conn            *gorm.DB
	TokenCollection *mongo.Collection
}

type GetsProductRepository struct {
	TokenCollection *mongo.Collection
}

type DeleteProductRepository struct {
	TokenCollection *mongo.Collection
}

type UpdateProductRepository struct {
	TokenCollection *mongo.Collection
}
