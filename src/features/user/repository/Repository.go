package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type WithdrawalUserRepository struct {
	GormDB          *gorm.DB
	TokenCollection *mongo.Collection
}
