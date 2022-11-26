package repository

import "go.mongodb.org/mongo-driver/mongo"

type RegisterProductRepository struct {
	TokenCollection *mongo.Collection
}

type GetProductRepository struct {
	TokenCollection *mongo.Collection
}

type GetsProductRepository struct {
	TokenCollection *mongo.Collection
}
