package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"main/common/dbCommon/mysqlCommon"
	"main/common/errorCommon"
	_interface "main/features/user/usecase/interface"
)

func NewWithdrawalUserRepository(gormDB *gorm.DB, tokenCollection *mongo.Collection) _interface.IWithdrawalUserRepository {
	return &WithdrawalUserRepository{GormDB: gormDB, TokenCollection: tokenCollection}
}

func (w *WithdrawalUserRepository) WithdrawalUser(ctx context.Context, userID string) error {
	result := w.GormDB.WithContext(ctx).Model(&mysqlCommon.GormUser{}).Where("is_deleted = ? AND id = ?", false, userID).Update("is_deleted", true)
	if result.Error != nil {
		return errorCommon.ErrorMsg(errorCommon.ErrInternalDB, errorCommon.Trace(), result.Error.Error(), errorCommon.ErrFromMysqlDB)
	}
	return nil
}

func (w *WithdrawalUserRepository) WithdrawalUserAuth(ctx context.Context, userID string) error {
	// Update with conditions and model value
	result := w.GormDB.WithContext(ctx).Model(&mysqlCommon.GormUserAuth{}).Where("is_deleted = ? AND user_id = ?", false, userID).Update("is_deleted", true)
	if result.Error != nil {
		return errorCommon.ErrorMsg(errorCommon.ErrInternalDB, errorCommon.Trace(), result.Error.Error(), errorCommon.ErrFromMysqlDB)
	}

	return nil
}
