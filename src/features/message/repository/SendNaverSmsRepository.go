package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"main/common/dbCommon/mongodbCommon"
	"main/common/errorCommon"
	_interface "main/features/message/usecase/interface"
	"main/features/product/domain"
)

func NewSendNaverSmsRepository(gormDB *gorm.DB, eventCollection *mongo.Collection) _interface.ISendNaverSmsRepository {
	return &SendNaverSmsRepository{GormDB: gormDB, EventCollection: eventCollection}
}

func (s *SendNaverSmsRepository) SaveNaverSmsEvent(ctx context.Context, msgEvent mongodbCommon.NaverSmsEventDTO) error {
	_, err := s.EventCollection.InsertOne(ctx, msgEvent)
	if err != nil {
		return errorCommon.ErrorMsg(errorCommon.ErrInternalDB, errorCommon.Trace(), domain.ErrInternalServerError, errorCommon.ErrFromMongoDB)
	}
	return nil
}
