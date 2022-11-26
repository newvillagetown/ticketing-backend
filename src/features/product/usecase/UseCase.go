package usecase

import (
	"main/common/dbCommon/mysqlCommon"
	"main/features/product/model/request"
	"main/features/product/model/response"
)

func ConvertToRegisterProductDTO(req request.ReqRegisterProduct) mysqlCommon.Product {
	id := mysqlCommon.PKIDGenerate()
	now := mysqlCommon.NowDateGenerate()
	result := mysqlCommon.Product{
		ID:          id,
		Created:     now,
		LastUpdated: now,
		IsDeleted:   false,
		Name:        req.Name,
		Description: req.Description,
		Category:    req.Category,
		PerAmount:   req.PerAmount,
		TotalCount:  req.TotalCount,
		RestCount:   req.RestCount,
		StartDate:   mysqlCommon.EpochToTimeString(req.StartDate),
		EndDate:     mysqlCommon.EpochToTimeString(req.EndDate),
	}
	if req.Image != nil {
		// s3 signed url
	}

	return result
}

func ConvertGetProductToRes(productDTO mysqlCommon.Product) response.ResGetProduct {
	result := response.ResGetProduct{
		ID:          productDTO.ID,
		Name:        productDTO.Name,
		Description: productDTO.Description,
		Category:    productDTO.Category,
		PerAmount:   productDTO.PerAmount,
		TotalCount:  productDTO.TotalCount,
		RestCount:   productDTO.RestCount,
		StartDate:   mysqlCommon.TimeStringToEpoch(productDTO.StartDate),
		EndDate:     mysqlCommon.TimeStringToEpoch(productDTO.EndDate),
	}
	return result
}
