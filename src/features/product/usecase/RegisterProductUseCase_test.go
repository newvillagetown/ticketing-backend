package usecase

import (
	"github.com/bxcodec/faker"
	"main/common/dbCommon/mysqlCommon"
)

func (s *UseCaseSuite) Test_usecase_Register() {

	//given
	var productDTO mysqlCommon.GormProduct
	err := faker.FakeData(&productDTO)
	s.NoError(err)

	s.RegisterProductRepo.On("CreateProduct", s.ctx, productDTO).Return(nil).Once()

	//when
	err = s.RegisterProductRepo.CreateProduct(s.ctx, productDTO)

	//then
	s.NoError(err)
}
