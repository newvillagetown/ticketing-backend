package repository

import "go.mongodb.org/mongo-driver/mongo"

type SignInGoogleOAuthRepository struct {
	TokenCollection *mongo.Collection
}
type SignOutGoogleOAuthRepository struct {
	TokenCollection *mongo.Collection
}
type CallbackGoogleOAuthRepository struct {
	TokenCollection *mongo.Collection
}
