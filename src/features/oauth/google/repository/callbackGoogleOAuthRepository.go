package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"main/common/dbCommon/mongodb"
	"main/common/oauthCommon/google"
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
	_, err := cc.TokenCollection.InsertOne(ctx, token)
	if err != nil {
		return err
	}
	return nil
}

func (cc *CallbackGoogleOAuthRepository) DeleteAllRefreshToken(authUser google.User) error {
	ctx := context.TODO()
	findData := bson.D{{"email", authUser.Email}}
	result, err := cc.TokenCollection.DeleteMany(ctx, findData)
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}
	fmt.Println(result.DeletedCount)
	return nil
}
