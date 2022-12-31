package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type SignInGoogleOAuthRepository struct {
	GormDB          *gorm.DB
	TokenCollection *mongo.Collection
}
type SignOutGoogleOAuthRepository struct {
	GormDB          *gorm.DB
	TokenCollection *mongo.Collection
}
type CallbackGoogleOAuthRepository struct {
	GormDB          *gorm.DB
	TokenCollection *mongo.Collection
}
