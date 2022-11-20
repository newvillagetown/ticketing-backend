package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"main/common/dbCommon/mongodbCommon"
	"main/common/dbCommon/mysqlCommon"
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

func (cc *CallbackGoogleOAuthRepository) FindOneUser(authUser google.User) (bool, error) {
	var email string
	err := mysqlCommon.MysqlDB.QueryRow("SELECT email FROM user WHERE email = ?", authUser.Email).Scan(&email)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			fmt.Println(err)
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (cc *CallbackGoogleOAuthRepository) CreateUser(userDTO mysqlCommon.User) error {
	// INSERT 문 실행
	fmt.Println("userDTO===================")
	fmt.Println(userDTO)
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
	fmt.Println("userAuthDTO=================")
	fmt.Println(userAuthDTO.Created)
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
