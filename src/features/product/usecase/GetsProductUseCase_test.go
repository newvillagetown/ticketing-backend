package usecase

import (
	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/require"
	"main/common/dbCommon/mysqlCommon"
)

func (s *UseCaseSuite) Test_usecase_Gets() {

	//given
	productDTOs := make([]mysqlCommon.GormProduct, 2, 2)
	for i := 0; i < 2; i++ {
		err := faker.FakeData(&productDTOs[i])
		s.NoError(err)
	}

	s.GetsProductRepo.On("FindProduct", s.ctx).Return(productDTOs, nil).Once()

	//when
	result, err := s.GetsProductRepo.FindProduct(s.ctx)

	//then
	s.NoError(err)
	require.Equal(s.T(), result, productDTOs)
}
