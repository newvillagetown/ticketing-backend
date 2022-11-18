package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"main/common/dbCommon/mongodb"
	_interface "main/features/oauth/google/usecase/interface"
)

func NewCallbackGoogleOAuthRepository(tokenCollection *mongo.Collection) _interface.ICallbackGoogleOAuthRepository {
	return &CallbackGoogleOAuthRepository{TokenCollection: tokenCollection}
}

func (cc *CallbackGoogleOAuthRepository) CallbackGoogle() error {

	return nil
}

func (cc *CallbackGoogleOAuthRepository) CreateRefreshToken(token mongodb.RefreshToken) error {
	ctx := context.TODO()
	fmt.Println("???")
	_, err := cc.TokenCollection.InsertOne(ctx, token)
	if err != nil {
		return err
	}
	return nil
}
