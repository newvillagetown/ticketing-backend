package usecase

import (
	"main/common/dbCommon/mysqlCommon"
	"main/features/product/model/request"
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
