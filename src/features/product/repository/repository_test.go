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
	_interface "main/features/product/usecase/interface"
	"testing"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	GetRepository      _interface.IGetProductRepository
	DeleteRepository   _interface.IDeleteProductRepository
	GetsRepository     _interface.IGetsProductRepository
	RegisterRepository _interface.IRegisterProductRepository
	UpdateRepository   _interface.IUpdateProductRepository
	product            *mysqlCommon.GormProduct
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
	s.GetRepository = NewGetProductRepository(s.DB, tokenCollection)
	s.GetsRepository = NewGetsProductRepository(s.DB, tokenCollection)
	//s.UpdateRepository = NewUpdateProductRepository(s.DB, tokenCollection)
	//s.RegisterRepository = NewRegisterProductRepository(s.DB, tokenCollection)
}
