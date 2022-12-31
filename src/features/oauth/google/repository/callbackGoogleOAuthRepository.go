package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"main/common/dbCommon/mongodbCommon"
	"main/common/dbCommon/mysqlCommon"
	"main/common/errorCommon"
	_interface "main/features/oauth/google/usecase/interface"
	"main/features/product/domain"
)

func NewCallbackGoogleOAuthRepository(gormDB *gorm.DB, tokenCollection *mongo.Collection) _interface.ICallbackGoogleOAuthRepository {
	return &CallbackGoogleOAuthRepository{GormDB: gormDB, TokenCollection: tokenCollection}
}

func (cc *CallbackGoogleOAuthRepository) CallbackGoogle(ctx context.Context) error {
	return nil
}

func (cc *CallbackGoogleOAuthRepository) CreateRefreshToken(ctx context.Context, token mongodbCommon.RefreshToken) error {
	_, err := cc.TokenCollection.InsertOne(ctx, token)
	if err != nil {
		return errorCommon.ErrorMsg(errorCommon.ErrInternalDB, errorCommon.Trace(), err.Error(), errorCommon.ErrFromMongoDB)
	}
	return nil
}

func (cc *CallbackGoogleOAuthRepository) DeleteAllRefreshToken(ctx context.Context, userDTO mysqlCommon.GormUser) error {
	findData := bson.D{{"email", userDTO.Email}}
	result, err := cc.TokenCollection.DeleteMany(ctx, findData)
	if err != nil && err != mongo.ErrNoDocuments {
		return errorCommon.ErrorMsg(errorCommon.ErrInternalDB, errorCommon.Trace(), err.Error(), errorCommon.ErrFromMongoDB)
	}
	fmt.Println(result.DeletedCount)
	return nil
}

func (cc *CallbackGoogleOAuthRepository) FindOneUser(ctx context.Context, userDTO mysqlCommon.GormUser) (mysqlCommon.GormUser, error) {
	result := cc.GormDB.Model(&mysqlCommon.GormUser{}).WithContext(ctx).Where("email = ? and is_deleted = ?", userDTO.Email, false).Scan(&userDTO)
	if result.RowsAffected == 0 {
		return mysqlCommon.GormUser{}, nil
	}
	if result.Error != nil {
		return mysqlCommon.GormUser{}, errorCommon.ErrorMsg(errorCommon.ErrInternalDB, errorCommon.Trace(), result.Error.Error(), errorCommon.ErrFromMysqlDB)
	}

	return userDTO, nil
}

func (cc *CallbackGoogleOAuthRepository) CreateUser(ctx context.Context, userDTO mysqlCommon.GormUser) error {
	result := cc.GormDB.WithContext(ctx).Create(&userDTO)
	if result.RowsAffected == 0 {
		return errorCommon.ErrorMsg(errorCommon.ErrBadParameter, errorCommon.Trace(), domain.ErrBadParamInput, errorCommon.ErrFromClient)
	}
	if result.Error != nil {
		return errorCommon.ErrorMsg(errorCommon.ErrInternalDB, errorCommon.Trace(), result.Error.Error(), errorCommon.ErrFromMysqlDB)
	}
	return nil
}
func (cc *CallbackGoogleOAuthRepository) CreateUserAuth(ctx context.Context, userAuthDTO mysqlCommon.GormUserAuth) error {
	result := cc.GormDB.WithContext(ctx).Create(&userAuthDTO)
	if result.RowsAffected == 0 {
		return errorCommon.ErrorMsg(errorCommon.ErrBadParameter, errorCommon.Trace(), domain.ErrBadParamInput, errorCommon.ErrFromClient)
	}
	if result.Error != nil {
		return errorCommon.ErrorMsg(errorCommon.ErrInternalDB, errorCommon.Trace(), result.Error.Error(), errorCommon.ErrFromMysqlDB)
	}
	return nil
}
