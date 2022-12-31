package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"main/common/errorCommon"
	_interface "main/features/oauth/google/usecase/interface"
)

func NewSignOutGoogleOAuthRepository(gormDB *gorm.DB, tokenCollection *mongo.Collection) _interface.ISignOutGoogleOAuthRepository {
	return &SignOutGoogleOAuthRepository{TokenCollection: tokenCollection}
}

func (s *SignOutGoogleOAuthRepository) SignOutGoogle(ctx context.Context) error {

	return nil
}
func (s *SignOutGoogleOAuthRepository) DeleteRefreshToken(ctx context.Context, email string) error {
	result, err := s.TokenCollection.DeleteMany(context.TODO(), bson.D{{"email", email}})
	if err != nil {
		return errorCommon.ErrorMsg(errorCommon.ErrInternalDB, errorCommon.Trace(), err.Error(), errorCommon.ErrFromMongoDB)
	}
	if result.DeletedCount >= 1 {
		return nil
	} else {
		return fmt.Errorf("invalid refresh token")
	}
	return nil
}
