package repository

import (
	"context"
	"fibonacci-spiral-matrix-go/internal/core/domain/user"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Insert(ctx context.Context, user *user.User) *gorm.DB
	GetByUsername(ctx context.Context, username string, user *user.User) *gorm.DB
	GetByEmail(ctx context.Context, email string, user *user.User) *gorm.DB
}
