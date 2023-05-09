package handler

import (
	"github.com/labstack/echo/v4"
	"main/features/system/test/repository"
	"main/features/system/test/usecase"
	"time"
)

func NewTestHandler(c *echo.Echo) {
	NewAuthTestHandler(c, usecase.NewAuthTestUseCase(repository.NewAuthTestRepository(), 8*time.Second))
}
