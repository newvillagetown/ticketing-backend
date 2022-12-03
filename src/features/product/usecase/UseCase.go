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

func ConvertUpdateProductNewProductDTO(req request.ReqUpdateProduct, productDTO mysqlCommon.Product) mysqlCommon.Product {
	now := mysqlCommon.NowDateGenerate()
	result := mysqlCommon.Product{
		ID:          productDTO.ID,
		Created:     productDTO.Created,
		LastUpdated: now,
		IsDeleted:   productDTO.IsDeleted,
		Name:        productDTO.Name,
		Description: productDTO.Description,
		Category:    productDTO.Category,
		PerAmount:   productDTO.PerAmount,
		ImgUrl:      productDTO.ImgUrl,
		TotalCount:  productDTO.TotalCount,
		RestCount:   productDTO.RestCount,
		StartDate:   productDTO.StartDate,
		EndDate:     productDTO.EndDate,
	}
	//true이면 변경할 데이터가 존재한다는 의미
	if NilCheckString(req.Name) {
		result.Name = req.Name
	}
	if NilCheckString(req.Description) {
		result.Description = req.Description
	}
	if NilCheckString(req.Category) {
		result.Category = req.Category
	}
	if NilCheckInt64(req.PerAmount) {
		result.PerAmount = req.PerAmount
	}
	if NilCheckInt64(req.TotalCount) {
		result.TotalCount = req.TotalCount
	}
	if NilCheckInt64(req.RestCount) {
		result.RestCount = req.RestCount
	}
	if NilCheckInt64(req.StartDate) {
		result.StartDate = mysqlCommon.EpochToTimeString(req.StartDate)
	}
	if NilCheckInt64(req.EndDate) {
		result.EndDate = mysqlCommon.EpochToTimeString(req.EndDate)
	}

	return result
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
