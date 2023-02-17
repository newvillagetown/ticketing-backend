package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type SendNaverSmsRepository struct {
	GormDB          *gorm.DB
	EventCollection *mongo.Collection
}
