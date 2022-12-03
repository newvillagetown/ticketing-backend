package usecase

import (
	"fmt"
	"main/common/dbCommon/mysqlCommon"
	"main/common/s3Common"
	"main/features/product/model/request"
	"main/features/product/model/response"
)

func ConvertToRegisterProductDTO(req request.ReqRegisterProduct) mysqlCommon.GormProduct {
	result := mysqlCommon.GormProduct{
		GormModel: mysqlCommon.GormModel{
			ID: mysqlCommon.PKIDGenerate(),
		},
		Name:        req.Name,
		Description: req.Description,
		Category:    req.Category,
		PerAmount:   req.PerAmount,
		TotalCount:  req.TotalCount,
		RestCount:   req.RestCount,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
	}

	if req.Image != nil {
		// s3 signed url
		filename, err := s3Common.ImageUpload(req.Image, s3Common.ImgTypeProduct)
		if err != nil {
			fmt.Println(err)
		}
		result.ImgUrl = filename
	}

	return result
}

func ConvertGetProductToRes(productDTO mysqlCommon.GormProduct) response.ResGetProduct {
	result := response.ResGetProduct{
		ID:          productDTO.GormModel.ID,
		Name:        productDTO.Name,
		Description: productDTO.Description,
		Category:    productDTO.Category,
		PerAmount:   productDTO.PerAmount,
		TotalCount:  productDTO.TotalCount,
		RestCount:   productDTO.RestCount,
		StartDate:   productDTO.StartDate,
		EndDate:     productDTO.EndDate,
	}
	signedURL, err := s3Common.ImageGetSignedURL(productDTO.ImgUrl, s3Common.ImgTypeProduct)
	if err != nil {
		fmt.Println(err)
	}
	result.Image = signedURL
	return result
}

func ConvertGetsProductToRes(productList []mysqlCommon.GormProduct) response.ResGetsProduct {
	result := response.ResGetsProduct{
		Count: int64(len(productList)),
	}
	arr := make([]response.GetsProduct, 0, len(productList))
	for i := 0; i < len(productList); i++ {
		cur := response.GetsProduct{
			ID:          productList[i].GormModel.ID,
			Name:        productList[i].Name,
			Description: productList[i].Description,
			Category:    productList[i].Category,
			PerAmount:   productList[i].PerAmount,
			TotalCount:  productList[i].TotalCount,
			RestCount:   productList[i].RestCount,
			StartDate:   productList[i].StartDate,
			EndDate:     productList[i].EndDate,
		}
		if productList[i].ImgUrl != "" {
			signedURL, err := s3Common.ImageGetSignedURL(productList[i].ImgUrl, s3Common.ImgTypeProduct)
			if err != nil {
				fmt.Println(err)
			}
			cur.Image = signedURL
		}
		arr = append(arr, cur)
	}
	result.Products = arr

	return result
}

func ConvertUpdateProductNewProductDTO(req request.ReqUpdateProduct, productDTO mysqlCommon.GormProduct) mysqlCommon.GormProduct {

	//true이면 변경할 데이터가 존재한다는 의미
	if NilCheckString(req.Name) {
		productDTO.Name = req.Name
	}
	if NilCheckString(req.Description) {
		productDTO.Description = req.Description
	}
	if NilCheckString(req.Category) {
		productDTO.Category = req.Category
	}
	if NilCheckInt64(req.PerAmount) {
		productDTO.PerAmount = req.PerAmount
	}
	if NilCheckInt64(req.TotalCount) {
		productDTO.TotalCount = req.TotalCount
	}
	if NilCheckInt64(req.RestCount) {
		productDTO.RestCount = req.RestCount
	}
	if NilCheckInt64(req.StartDate) {
		productDTO.StartDate = req.StartDate
	}
	if NilCheckInt64(req.EndDate) {
		productDTO.EndDate = req.EndDate
	}

	return productDTO
}

func NilCheckString(str string) bool {
	if str != "" {
		return true
	}
	return false
}

func NilCheckInt64(num int64) bool {
	if num != 0 {
		return true
	}
	return false
}
