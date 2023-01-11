package usecase

import (
	"github.com/bxcodec/faker"
	"main/features/product/domain/request"
)

func (s *UseCaseSuite) Test_usecase_Delete() {

	//given
	var req request.ReqDeleteProduct
	err := faker.FakeData(&req)
	s.NoError(err)

	s.DeleteProductRepo.On("FindOneAndDeleteUpdateProduct", s.ctx, req.ProductID).Return(nil).Once()

	//when
	err = s.DeleteProductRepo.FindOneAndDeleteUpdateProduct(s.ctx, req.ProductID)

	//then
	s.NoError(err)
}
