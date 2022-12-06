package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/bxcodec/faker"
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

	repository _interface.IGetProductRepository
	product    *mysqlCommon.GormProduct
}

type repo struct {
	DB *gorm.DB
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
	s.repository = NewGetProductRepository(s.DB, tokenCollection)
}

func (s *Suite) Test_repository_Get() {
	var mockDBProduct mysqlCommon.GormProduct
	err := faker.FakeData(&mockDBProduct)

	s.NoError(err)
	rows := sqlmock.NewRows([]string{"name", "description", "category", "per_amount", "img_url", "total_count", "rest_count", "start_date", "end_date"}).
		AddRow(mockDBProduct.Name, mockDBProduct.Description, mockDBProduct.Category, mockDBProduct.PerAmount, mockDBProduct.ImgUrl, mockDBProduct.TotalCount, mockDBProduct.RestCount, mockDBProduct.StartDate, mockDBProduct.EndDate)
	s.mock.ExpectQuery("SELECT").WillReturnRows(rows)

	ctx := context.TODO()
	res, err := s.repository.FindOneProduct(ctx, mockDBProduct.GormModel.ID)
	fmt.Println(err)
	require.NoError(s.T(), err)
	require.Equal(s.T(), mockDBProduct.Name, res.Name)
}
