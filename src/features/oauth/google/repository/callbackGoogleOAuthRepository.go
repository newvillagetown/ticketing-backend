package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"main/common/dbCommon/mongodbCommon"
	"main/common/dbCommon/mysqlCommon"
	"main/common/errorCommon"
	"main/common/oauthCommon/google"
	_interface "main/features/oauth/google/usecase/interface"
)

func NewCallbackGoogleOAuthRepository(tokenCollection *mongo.Collection) _interface.ICallbackGoogleOAuthRepository {
	return &CallbackGoogleOAuthRepository{TokenCollection: tokenCollection}
}

func (cc *CallbackGoogleOAuthRepository) CallbackGoogle() error {
	return nil
}

func (cc *CallbackGoogleOAuthRepository) CreateRefreshToken(token mongodbCommon.RefreshToken) error {
	ctx := context.TODO()
	_, err := cc.TokenCollection.InsertOne(ctx, token)
	if err != nil {
		return errorCommon.ErrorMsg(errorCommon.ErrInternalDB, errorCommon.Trace(), err.Error(), errorCommon.ErrFromMongoDB)
	}
	return nil
}

func (cc *CallbackGoogleOAuthRepository) DeleteAllRefreshToken(authUser google.User) error {
	ctx := context.TODO()
	findData := bson.D{{"email", authUser.Email}}
	result, err := cc.TokenCollection.DeleteMany(ctx, findData)
	if err != nil && err != mongo.ErrNoDocuments {
		return errorCommon.ErrorMsg(errorCommon.ErrInternalDB, errorCommon.Trace(), err.Error(), errorCommon.ErrFromMongoDB)
	}
	fmt.Println(result.DeletedCount)
	return nil
}

func (cc *CallbackGoogleOAuthRepository) FindOneUser(authUser google.User) (string, error) {
	var id string
	err := mysqlCommon.MysqlDB.QueryRow("SELECT id FROM user WHERE email = ?", authUser.Email).Scan(&id)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			fmt.Println(err)
			return "", err
		}
		return id, nil
	}
	return id, nil
}

func (cc *CallbackGoogleOAuthRepository) CreateUser(userDTO mysqlCommon.User) error {
	// INSERT 문 실행
	result, err := mysqlCommon.MysqlDB.Exec("INSERT INTO user VALUES (?, ?,?,?,?)", userDTO.ID, userDTO.Name, userDTO.Email, userDTO.Created, userDTO.IsDeleted)
	if err != nil {
		fmt.Println(err)
		return err
	}
	n, err := result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row inserted.")
	}
	return nil
}
func (cc *CallbackGoogleOAuthRepository) CreateUserAuth(userAuthDTO mysqlCommon.UserAuth) error {
	result, err := mysqlCommon.MysqlDB.Exec("INSERT INTO userauth VALUES (?, ?,?,?,?,?)", userAuthDTO.ID, userAuthDTO.Provider, userAuthDTO.UserID, userAuthDTO.LastSignIn, userAuthDTO.Created, userAuthDTO.IsDeleted)
	if err != nil {
		fmt.Println(err)
		return err
	}
	n, err := result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row inserted.")
	}
	return nil
}
