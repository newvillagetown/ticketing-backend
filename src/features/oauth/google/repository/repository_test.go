package repository

import (
	"database/sql"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/DATA-DOG/go-sqlmock.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main/common/dbCommon/mysqlCommon"
	_interface "main/features/oauth/google/usecase/interface"
	"testing"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	CallbackGoogleOAuthRepository _interface.ICallbackGoogleOAuthRepository
	product                       *mysqlCommon.GormProduct
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)
	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)
	s.DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	require.NoError(s.T(), err)
	var tokenCollection *mongo.Collection
	s.CallbackGoogleOAuthRepository = NewCallbackGoogleOAuthRepository(s.DB, tokenCollection)
}
