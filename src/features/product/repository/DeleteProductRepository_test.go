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

func (s *RepositorySuite) Test_repository_FindOneAndDeleteUpdateProduct() {
	var mockDBProduct mysqlCommon.GormProduct
	err := faker.FakeData(&mockDBProduct)
	if err != nil {
		fmt.Println(err)
	}
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `gorm_products` SET `is_deleted`=?,`updated_at`=? WHERE id = ?")).WithArgs(true, time.Now().Unix(), mockDBProduct.GormModel.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	ctx := context.TODO()
	err = s.DeleteRepository.FindOneAndDeleteUpdateProduct(ctx, mockDBProduct.GormModel.ID)
	require.NoError(s.T(), err)
}
