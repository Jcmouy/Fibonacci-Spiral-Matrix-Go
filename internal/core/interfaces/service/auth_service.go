package service

import (
	"context"
	"fibonacci-spiral-matrix-go/internal/api/dto"
	"fibonacci-spiral-matrix-go/internal/auth"
	"fibonacci-spiral-matrix-go/internal/core/domain/user"
	"gorm.io/gorm"
)

type AuthService interface {
	HashUserPassword(context.Context, *dto.SignupRequestDto) error
	CreateUserRecord(context.Context, *user.User) error
	VerifyUserEmail(context.Context, string, *user.User) *gorm.DB
	CheckUsername(context.Context, *user.User, string) *gorm.DB
	CheckUserPassword(context.Context, user.User, string) error
	CreateJwtWrapper(context.Context) auth.JwtWrapper
	GenerateToken(context.Context, auth.JwtWrapper, string, string, string) (string, error)
}
