package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	_interface "main/features/user/usecase/interface"
)

func NewWithdrawalUserRepository(tokenCollection *mongo.Collection) _interface.IWithdrawalUserRepository {
	return &WithdrawalUserRepository{TokenCollection: tokenCollection}
}

func (w *WithdrawalUserRepository) WithdrawalUser() error {

	return nil
}
