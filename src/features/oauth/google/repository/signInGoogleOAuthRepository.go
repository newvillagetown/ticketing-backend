package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	_interface "main/features/oauth/google/usecase/interface"
)

func NewSignInGoogleOAuthRepository(gormDB *gorm.DB, tokenCollection *mongo.Collection) _interface.ISignInGoogleOAuthRepository {
	return &SignInGoogleOAuthRepository{TokenCollection: tokenCollection}
}

func (s *SignInGoogleOAuthRepository) SignInGoogle(ctx context.Context) error {

	return nil
}
