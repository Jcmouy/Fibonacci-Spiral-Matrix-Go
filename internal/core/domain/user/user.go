package user

import (
	"fibonacci-spiral-matrix-go/internal/api/dto"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
	Role     string
}

func (user *User) FromModel(signupRequestDto dto.SignupRequestDto) User {
	user.Username = signupRequestDto.Username
	user.Email = signupRequestDto.Email
	user.Password = signupRequestDto.Password
	user.Role = signupRequestDto.Role
	return *user
}
