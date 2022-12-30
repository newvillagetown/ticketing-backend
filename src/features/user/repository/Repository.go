package repository

import "go.mongodb.org/mongo-driver/mongo"

type WithdrawalUserRepository struct {
	TokenCollection *mongo.Collection
}
