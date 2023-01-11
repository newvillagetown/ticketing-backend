package usecase

import (
	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/require"
	"main/common/dbCommon/mysqlCommon"
	"main/features/product/domain/request"
)

func (s *UseCaseSuite) Test_usecase_Get() {

	//given
	var req request.ReqGetProduct
	err := faker.FakeData(&req)
	s.NoError(err)
	var productDTO mysqlCommon.GormProduct
	err = faker.FakeData(&productDTO)
	s.NoError(err)
	s.repo.On("FindOneProduct", s.ctx, req.ProductID).Return(productDTO, nil).Once()

	//when
	result, err := s.repo.FindOneProduct(s.ctx, req.ProductID)

	//then
	require.Equal(s.T(), result, productDTO)
}
