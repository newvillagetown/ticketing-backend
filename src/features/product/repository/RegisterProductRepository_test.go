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

func (s *Suite) Test_repository_CreateProduct() {
	var mockDBProduct mysqlCommon.GormProduct
	err := faker.FakeData(&mockDBProduct)
	if err != nil {
		fmt.Println(err)
	}
	mockDBProduct.GormModel.CreatedAt = time.Now().Unix()
	mockDBProduct.GormModel.UpdatedAt = time.Now().Unix()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `gorm_products` (`id`,`created_at`,`updated_at`,`is_deleted`,`name`,`description`,`category`,`per_amount`,`img_url`,`total_count`,`rest_count`,`start_date`,`end_date`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)")).
		WithArgs(mockDBProduct.GormModel.ID, mockDBProduct.GormModel.CreatedAt, mockDBProduct.GormModel.UpdatedAt, mockDBProduct.GormModel.IsDeleted, mockDBProduct.Name, mockDBProduct.Description, mockDBProduct.Category, mockDBProduct.PerAmount, mockDBProduct.ImgUrl, mockDBProduct.TotalCount, mockDBProduct.RestCount, mockDBProduct.StartDate, mockDBProduct.EndDate).WillReturnResult(sqlmock.NewResult(1, 1))
	ctx := context.TODO()
	err = s.RegisterRepository.CreateProduct(ctx, mockDBProduct)
	require.NoError(s.T(), err)
}
