package repository

import (
	"context"
	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/require"
	"gopkg.in/DATA-DOG/go-sqlmock.v2"
	"main/common/dbCommon/mysqlCommon"
	"regexp"
)

func (s *Suite) Test_repository_FindOneUser() {
	var mockDBUser mysqlCommon.GormUser
	err := faker.FakeData(&mockDBUser)
	s.NoError(err)
	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "is_deleted", "name", "email"}).
		AddRow(mockDBUser.GormModel.ID, mockDBUser.GormModel.CreatedAt, mockDBUser.GormModel.UpdatedAt, mockDBUser.GormModel.IsDeleted, mockDBUser.Name, mockDBUser.Email)
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `gorm_users` WHERE email = ? and is_deleted")).WithArgs(mockDBUser.Email, false).WillReturnRows(rows)
	ctx := context.TODO()
	_, err = s.CallbackGoogleOAuthRepository.FindOneUser(ctx, mockDBUser)
	require.NoError(s.T(), err)
}

func (s *Suite) Test_repository_CreateUser() {
	var mockDBUser mysqlCommon.GormUser
	err := faker.FakeData(&mockDBUser)

	s.NoError(err)

	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `gorm_users` (`id`,`created_at`,`updated_at`,`is_deleted`,`name`,`email`) VALUES (?,?,?,?,?,?)")).
		WithArgs(mockDBUser.GormModel.ID, mockDBUser.GormModel.CreatedAt, mockDBUser.GormModel.UpdatedAt, mockDBUser.GormModel.IsDeleted, mockDBUser.Name, mockDBUser.Email).WillReturnResult(sqlmock.NewResult(1, 1))
	ctx := context.TODO()
	err = s.CallbackGoogleOAuthRepository.CreateUser(ctx, mockDBUser)
	require.NoError(s.T(), err)
}

func (s *Suite) Test_repository_CreateUserAuth() {
	var mockDBUserAuth mysqlCommon.GormUserAuth
	err := faker.FakeData(&mockDBUserAuth)

	s.NoError(err)

	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `gorm_user_auths` (`id`,`created_at`,`updated_at`,`is_deleted`,`provider`,`user_id`,`last_sign_in`) VALUES (?,?,?,?,?,?,?)")).
		WithArgs(mockDBUserAuth.GormModel.ID, mockDBUserAuth.GormModel.CreatedAt, mockDBUserAuth.GormModel.UpdatedAt, mockDBUserAuth.GormModel.IsDeleted, mockDBUserAuth.Provider, mockDBUserAuth.UserID, mockDBUserAuth.LastSignIn).WillReturnResult(sqlmock.NewResult(1, 1))
	ctx := context.TODO()
	err = s.CallbackGoogleOAuthRepository.CreateUserAuth(ctx, mockDBUserAuth)
	require.NoError(s.T(), err)
}
