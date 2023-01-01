package repository

import (
	"context"
	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/require"
	"gopkg.in/DATA-DOG/go-sqlmock.v2"
	"main/common/dbCommon/mysqlCommon"
	"regexp"
)

func (s *Suite) Test_repository_FindProduct() {
	mockProducts := make([]mysqlCommon.GormProduct, 2, 2)
	for i := 0; i < len(mockProducts); i++ {
		err := faker.FakeData(&mockProducts[i])
		s.NoError(err)
	}
	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "is_deleted", "name", "description", "category", "per_amount", "img_url", "total_count", "rest_count", "start_date", "end_date"}).
		AddRow(mockProducts[0].GormModel.ID, mockProducts[0].GormModel.CreatedAt, mockProducts[0].GormModel.UpdatedAt, mockProducts[0].GormModel.IsDeleted, mockProducts[0].Name, mockProducts[0].Description, mockProducts[0].Category,
			mockProducts[0].PerAmount, mockProducts[0].ImgUrl, mockProducts[0].TotalCount, mockProducts[0].RestCount, mockProducts[0].StartDate, mockProducts[0].EndDate).
		AddRow(mockProducts[1].GormModel.ID, mockProducts[1].GormModel.CreatedAt, mockProducts[1].GormModel.UpdatedAt, mockProducts[1].GormModel.IsDeleted, mockProducts[1].Name, mockProducts[1].Description, mockProducts[1].Category,
			mockProducts[1].PerAmount, mockProducts[1].ImgUrl, mockProducts[1].TotalCount, mockProducts[1].RestCount, mockProducts[1].StartDate, mockProducts[1].EndDate)
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `gorm_products`")).WillReturnRows(rows)

	ctx := context.TODO()
	res, err := s.GetsRepository.FindProduct(ctx)
	require.NoError(s.T(), err)
	require.Equal(s.T(), mockProducts[0].Name, res[0].Name)
	require.Equal(s.T(), mockProducts[1].Name, res[1].Name)
}
