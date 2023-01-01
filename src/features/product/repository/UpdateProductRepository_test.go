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

func (s *Suite) Test_repository_update_FindOneProduct() {
	var mockDBProduct mysqlCommon.GormProduct
	err := faker.FakeData(&mockDBProduct)
	if err != nil {
		fmt.Println(err)
	}
	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "is_deleted", "name", "description", "category", "per_amount", "img_url", "total_count", "rest_count", "start_date", "end_date"}).
		AddRow(mockDBProduct.GormModel.ID, mockDBProduct.GormModel.CreatedAt, mockDBProduct.GormModel.UpdatedAt, mockDBProduct.GormModel.IsDeleted, mockDBProduct.Name, mockDBProduct.Description, mockDBProduct.Category, mockDBProduct.PerAmount, mockDBProduct.ImgUrl, mockDBProduct.TotalCount, mockDBProduct.RestCount, mockDBProduct.StartDate, mockDBProduct.EndDate)
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `gorm_products` WHERE id = ?")).WithArgs(mockDBProduct.GormModel.ID).WillReturnRows(rows)
	ctx := context.TODO()
	res, err := s.UpdateRepository.FindOneProduct(ctx, mockDBProduct.GormModel.ID)
	require.NoError(s.T(), err)
	require.Equal(s.T(), mockDBProduct.Name, res.Name)
}

func (s *Suite) Test_repository_update_FindOneAndUpdateProduct() {
	var mockDBProduct mysqlCommon.GormProduct
	err := faker.FakeData(&mockDBProduct)
	if err != nil {
		fmt.Println(err)
	}
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `gorm_products` SET `created_at`=?,`updated_at`=?,`is_deleted`=?,`name`=?,`description`=?,`category`=?,`per_amount`=?,`img_url`=?,`total_count`=?,`rest_count`=?,`start_date`=?,`end_date`=? WHERE `id` = ?")).
		WithArgs(mockDBProduct.GormModel.CreatedAt, time.Now().Unix(), mockDBProduct.GormModel.IsDeleted, mockDBProduct.Name, mockDBProduct.Description, mockDBProduct.Category, mockDBProduct.PerAmount, mockDBProduct.ImgUrl, mockDBProduct.TotalCount, mockDBProduct.RestCount, mockDBProduct.StartDate, mockDBProduct.EndDate, mockDBProduct.GormModel.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	ctx := context.TODO()
	err = s.UpdateRepository.FindOneAndUpdateProduct(ctx, mockDBProduct)
	require.NoError(s.T(), err)
}
