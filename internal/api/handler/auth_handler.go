package handler

import (
	"fibonacci-spiral-matrix-go/internal/api/dto"
	"fibonacci-spiral-matrix-go/internal/api/httperror"
	"fibonacci-spiral-matrix-go/internal/core/domain/user"
	"fibonacci-spiral-matrix-go/internal/core/interfaces/service"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

type AuthHandler struct {
	AuthService service.AuthService
}

type LoginPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Jwt                 string `json:"jwt"`
	UserDetailsId       uint   `json:"userdetailsid"`
	UserDetailsUsername string `json:"userdetailsusername"`
	UserDetailsEmail    string `json:"userdetailsemail"`
	Roles               string `json:"roles"`
}

type SignupResponse struct {
	Message string `json:"message"`
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
	return AuthHandler{AuthService: authService}
}

func (authHan *AuthHandler) Register(apiRouteGroup *gin.RouterGroup) {
	router := apiRouteGroup.Group("/auth")
	{
		router.POST("/signup", authHan.SignUp)
		router.POST("/login", authHan.Login)
	}
}

// SignUp godoc
// @Summary SignUp a user
// @Description SignUp a user with field email, username, password
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body user.User true "Insert user"
// @Success 200 {object} dto.SignupRequestDto
// @Failure 400 {object} dto.ErrorOutput
// @Failure 404 {object} dto.ErrorOutput
// @Router /api/v1/auth/signup [post]
func (authHan *AuthHandler) SignUp(ctx *gin.Context) {
	span := opentracing.GlobalTracer().StartSpan("AuthHandler-SignUp")
	spanContext := opentracing.ContextWithSpan(ctx.Request.Context(), span)
	defer span.Finish()

	var signupRequestDto dto.SignupRequestDto

	err := ctx.ShouldBindJSON(&signupRequestDto)
	if err != nil {
		log.Error(errors.Wrap(err, "invalid json"))
		httperror.NewError(ctx, http.StatusBadRequest, "Invalid json", err)
		return
	}

	err = authHan.AuthService.HashUserPassword(spanContext, &signupRequestDto)
	if err != nil {
		log.Error(errors.Wrap(err, "error hashing password"))
		httperror.NewError(ctx, http.StatusInternalServerError, "Error hashing password", err)
		return
	}

	user := user.User{}
	user.FromModel(signupRequestDto)

	err = authHan.AuthService.CreateUserRecord(spanContext, &user)
	if err != nil {
		log.Error(errors.Wrap(err, "error creating user"))
		httperror.NewError(ctx, http.StatusInternalServerError, "Error creating user", err)
		return
	}

	signupResponse := SignupResponse{
		Message: "User registered successfully!",
	}

	ctx.JSON(http.StatusOK, signupResponse)
}

// Login godoc
// @Summary Login a user
// @Description Login a user by username and password
// @Tags login
// @Accept json
// @Produce json
// @Param payload body LoginPayload true "Login user"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} dto.ErrorOutput
// @Failure 404 {object} dto.ErrorOutput
// @Router /api/v1/auth/login [post]
func (authHan *AuthHandler) Login(ctx *gin.Context) {
	span := opentracing.GlobalTracer().StartSpan("AuthHandler-SignUp")
	spanContext := opentracing.ContextWithSpan(ctx.Request.Context(), span)
	defer span.Finish()

	var payload LoginPayload
	var userInput user.User

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(400, gin.H{
			"msg": "invalid json",
		})
		ctx.Abort()
		return
	}

	if payload.Email != "" {
		result := authHan.AuthService.VerifyUserEmail(spanContext, payload.Email, &userInput)

		if result.Error == gorm.ErrRecordNotFound {
			ctx.JSON(401, gin.H{
				"msg": "invalid user credentials",
			})
			ctx.Abort()
			return
		}
	}

	result := authHan.AuthService.CheckUsername(spanContext, &userInput, payload.Username)

	if result.Error == gorm.ErrRecordNotFound {
		ctx.JSON(401, gin.H{
			"msg": "invalid user credentials",
		})
		ctx.Abort()
		return
	}

	err = authHan.AuthService.CheckUserPassword(spanContext, userInput, payload.Password)
	if err != nil {
		log.Println(err)
		ctx.JSON(401, gin.H{
			"msg": "invalid user credentials",
		})
		ctx.Abort()
		return
	}

	jwtWrapper := authHan.AuthService.CreateJwtWrapper(spanContext)

	signedToken, err := authHan.AuthService.GenerateToken(spanContext, jwtWrapper, userInput.Email, userInput.Username, userInput.Role)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{
			"msg": "error signing token",
		})
		ctx.Abort()
		return
	}

	tokenResponse := LoginResponse{
		Jwt:                 signedToken,
		UserDetailsId:       userInput.ID,
		UserDetailsUsername: userInput.Username,
		UserDetailsEmail:    userInput.Email,
		Roles:               userInput.Role,
	}

	ctx.JSON(http.StatusOK, tokenResponse)

	return
}
