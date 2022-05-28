package dto

type SignupRequestDto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role" validate:"required,eq=ADMIN|eq=USER"`
}
