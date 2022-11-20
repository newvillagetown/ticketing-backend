package repository

import "go.mongodb.org/mongo-driver/mongo"

type SignInGoogleOAuthRepository struct {
}
type SignOutGoogleOAuthRepository struct {
}
type CallbackGoogleOAuthRepository struct {
	TokenCollection *mongo.Collection
}
