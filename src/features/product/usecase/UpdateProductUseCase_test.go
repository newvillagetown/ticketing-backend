package usecase

import (
	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/require"
	"main/common/dbCommon/mysqlCommon"
	"main/features/product/domain/request"
)

func (s *UseCaseSuite) Test_usecase_Update() {

	//given
	var req request.ReqUpdateProduct
	err := faker.FakeData(&req)
	s.NoError(err)
	var productDTO mysqlCommon.GormProduct
	err = faker.FakeData(&productDTO)
	s.NoError(err)
	productDTO.GormModel.ID = req.ProductID
	updatedProductDTO := ConvertUpdateProductNewProductDTO(req, productDTO)
	s.UpdateProductRepo.On("FindOneProduct", s.ctx, req.ProductID).Return(productDTO, nil).Once()
	s.UpdateProductRepo.On("FindOneAndUpdateProduct", s.ctx, updatedProductDTO).Return(nil).Once()

	//when
	mockProductDTO, err := s.UpdateProductRepo.FindOneProduct(s.ctx, req.ProductID)
	mockUpdatedProductDTO := ConvertUpdateProductNewProductDTO(req, mockProductDTO)
	err = s.UpdateProductRepo.FindOneAndUpdateProduct(s.ctx, mockUpdatedProductDTO)
	//then
	s.NoError(err)
	require.Equal(s.T(), mockUpdatedProductDTO, updatedProductDTO)
}
