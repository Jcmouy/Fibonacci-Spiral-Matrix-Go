package services

import (
	"context"
	"fibonacci-spiral-matrix-go/internal/api/dto"
	"fibonacci-spiral-matrix-go/internal/auth"
	"fibonacci-spiral-matrix-go/internal/core/domain/user"
	"fibonacci-spiral-matrix-go/internal/core/interfaces/repository"
	"github.com/opentracing/opentracing-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	authRepository repository.AuthRepository
}

func NewAuthServiceService(authRepository repository.AuthRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		authRepository: authRepository,
	}
}

func (authSerImpl *AuthServiceImpl) HashUserPassword(ctx context.Context, signupRequestDto *dto.SignupRequestDto) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "HashUserPassword")
	defer span.Finish()

	bytes, err := bcrypt.GenerateFromPassword([]byte(signupRequestDto.Password), 14)
	if err != nil {
		return err
	}

	signupRequestDto.Password = string(bytes)

	return nil
}

func (authSerImpl AuthServiceImpl) CreateUserRecord(ctx context.Context, user *user.User) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "CreateUserRecord")
	defer span.Finish()

	result := authSerImpl.authRepository.Insert(ctx, user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (authSerImpl AuthServiceImpl) VerifyUserEmail(ctx context.Context, providedEmail string, user *user.User) *gorm.DB {
	span, _ := opentracing.StartSpanFromContext(ctx, "VerifyUserEmail")
	defer span.Finish()

	return authSerImpl.authRepository.GetByEmail(ctx, providedEmail, user)
}

func (authSerImpl AuthServiceImpl) CheckUsername(ctx context.Context, user *user.User, providedUsername string) *gorm.DB {
	span, _ := opentracing.StartSpanFromContext(ctx, "CheckUsername")
	defer span.Finish()

	return authSerImpl.authRepository.GetByUsername(ctx, providedUsername, user)
}

func (authSerImpl AuthServiceImpl) CheckUserPassword(ctx context.Context, user user.User, providedPassword string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "CheckUserPassword")
	defer span.Finish()

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}

func (authSerImpl AuthServiceImpl) CreateJwtWrapper(ctx context.Context) auth.JwtWrapper {
	span, _ := opentracing.StartSpanFromContext(ctx, "CreateJwtWrapper")
	defer span.Finish()

	return auth.JwtWrapper{
		SecretKey:       "verysecretkey",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}
}

func (authSerImpl AuthServiceImpl) GenerateToken(ctx context.Context, jwtWrapper auth.JwtWrapper, userEmail, userName, role string) (string, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "GenerateToken")
	defer span.Finish()

	return jwtWrapper.GenerateToken(userEmail, userName, role)
}
