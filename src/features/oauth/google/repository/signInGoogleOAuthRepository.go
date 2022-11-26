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

func NewSignInGoogleOAuthRepository(tokenCollection *mongo.Collection) _interface.ISignInGoogleOAuthRepository {
	return &SignInGoogleOAuthRepository{TokenCollection: tokenCollection}
}

func (s *SignInGoogleOAuthRepository) SignInGoogle() error {

	return nil
}

func (s *SignInGoogleOAuthRepository) CreateRefreshToken(token mongodbCommon.RefreshToken) error {
	ctx := context.TODO()
	_, err := s.TokenCollection.InsertOne(ctx, token)
	if err != nil {
		return err
	}
	return nil
}

func (s *SignInGoogleOAuthRepository) DeleteAllRefreshToken(authUser google.User) error {
	ctx := context.TODO()
	findData := bson.D{{"email", authUser.Email}}
	result, err := s.TokenCollection.DeleteMany(ctx, findData)
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}
	fmt.Println(result.DeletedCount)
	return nil
}

func (s *SignInGoogleOAuthRepository) FindOneUser(authUser google.User) (string, error) {
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

func (s *SignInGoogleOAuthRepository) CreateUser(userDTO mysqlCommon.User) error {
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
func (s *SignInGoogleOAuthRepository) CreateUserAuth(userAuthDTO mysqlCommon.UserAuth) error {
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
