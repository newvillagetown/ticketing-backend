package repository

import (
	_interface "main/features/system/test/usecase/interface"
)

func NewAuthTestRepository() _interface.IAuthTestRepository {
	return &AuthTestRepository{}
}
