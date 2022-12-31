package repository

import (
	"context"
	"fmt"
	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/require"
	"gopkg.in/DATA-DOG/go-sqlmock.v2"
	"main/common/dbCommon/mysqlCommon"
	"regexp"
	"time"
)

func (s *Suite) Test_repository_withdrawalUser() {
	var mockDBUser mysqlCommon.GormUser
	err := faker.FakeData(&mockDBUser)
	if err != nil {
		fmt.Println(err)
	}
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `gorm_users` SET `is_deleted`=?,`updated_at`=? WHERE is_deleted = ? AND id = ?")).WithArgs(true, time.Now().Unix(), false, mockDBUser.GormModel.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	ctx := context.TODO()
	err = s.WithdrawalUserRepository.WithdrawalUser(ctx, mockDBUser.GormModel.ID)
	require.NoError(s.T(), err)
}

func (s *Suite) Test_repository_withdrawalUserAuth() {
	var mockDBUser mysqlCommon.GormUser
	err := faker.FakeData(&mockDBUser)
	if err != nil {
		fmt.Println(err)
	}
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `gorm_user_auths` SET `is_deleted`=?,`updated_at`=? WHERE is_deleted = ? AND user_id = ?")).WithArgs(true, time.Now().Unix(), false, mockDBUser.GormModel.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	ctx := context.TODO()
	err = s.WithdrawalUserRepository.WithdrawalUserAuth(ctx, mockDBUser.GormModel.ID)
	require.NoError(s.T(), err)
}
