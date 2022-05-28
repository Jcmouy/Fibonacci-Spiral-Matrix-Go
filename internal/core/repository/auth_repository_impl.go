package repository

import (
	"context"
	"fibonacci-spiral-matrix-go/internal/config/database"
	"fibonacci-spiral-matrix-go/internal/core/domain/user"
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
}

func NewAuthRepository() *AuthRepositoryImpl {
	return &AuthRepositoryImpl{}
}

func (a AuthRepositoryImpl) Insert(ctx context.Context, user *user.User) *gorm.DB {
	span, _ := opentracing.StartSpanFromContext(ctx, "Insert")
	defer span.Finish()

	return database.GlobalDB.Create(user)
}

func (a AuthRepositoryImpl) GetByUsername(ctx context.Context, username string, user *user.User) *gorm.DB {
	span, _ := opentracing.StartSpanFromContext(ctx, "GetByUsername")
	defer span.Finish()

	return database.GlobalDB.Where("username = ?", username).First(user)
}

func (a AuthRepositoryImpl) GetByEmail(ctx context.Context, email string, user *user.User) *gorm.DB {
	span, _ := opentracing.StartSpanFromContext(ctx, "GetByEmail")
	defer span.Finish()

	return database.GlobalDB.Where("email = ?", email).First(user)
}
