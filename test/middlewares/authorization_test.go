package middlewares

import (
	"bytes"
	"encoding/json"
	"fibonacci-spiral-matrix-go/internal/api/dto"
	"fibonacci-spiral-matrix-go/internal/api/middlewares"
	"fibonacci-spiral-matrix-go/internal/auth"
	"fibonacci-spiral-matrix-go/internal/config/database"
	"fibonacci-spiral-matrix-go/internal/core/domain/user"
	"fibonacci-spiral-matrix-go/internal/core/services"
	"fibonacci-spiral-matrix-go/internal/wired"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthorizationNoHeader(t *testing.T) {
	router := gin.Default()
	router.Use(middlewares.Authorization())

	fiboSpiralMatrixHandler, _ := wired.InitializeFiboSpiralMatrixHandler()

	router.GET("/api/user/spiral", fiboSpiralMatrixHandler.GetSpiralMatrix)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/user/spiral", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 403, w.Code)
}

func TestAuthorizationInvalidTokenFormat(t *testing.T) {
	router := gin.Default()
	router.Use(middlewares.Authorization())

	fiboSpiralMatrixHandler, _ := wired.InitializeFiboSpiralMatrixHandler()

	router.GET("/api/user/spiral", fiboSpiralMatrixHandler.GetSpiralMatrix)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/user/spiral", nil)
	req.Header.Add("Authorization", "test")

	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestAuthorizationInvalidToken(t *testing.T) {
	invalidToken := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	router := gin.Default()
	router.Use(middlewares.Authorization())

	fiboSpiralMatrixHandler, _ := wired.InitializeFiboSpiralMatrixHandler()

	router.GET("/api/user/spiral", fiboSpiralMatrixHandler.GetSpiralMatrix)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/user/spiral", nil)
	req.Header.Add("Authorization", invalidToken)

	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
}

func TestValidToken(t *testing.T) {
	var response dto.FibonacciSpiralMatrixDto

	signupRequestDto := dto.SignupRequestDto{
		Email:    "test@email.com",
		Password: "secret",
		Username: "Test User",
		Role:     "USER",
	}

	request := dto.MatrixInput{
		Row:    "4",
		Column: "5",
	}

	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "verysecretkey",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	err := database.InitDatabase()
	assert.NoError(t, err)

	err = database.GlobalDB.AutoMigrate(&user.User{})
	assert.NoError(t, err)

	router := gin.Default()
	router.Use(middlewares.Authorization())

	jsonRequest, err := json.Marshal(&request)
	assert.NoError(t, err)

	token, err := jwtWrapper.GenerateToken(signupRequestDto.Email, signupRequestDto.Username, signupRequestDto.Role)
	assert.NoError(t, err)

	fiboSpiralMatrixHandler, _ := wired.InitializeFiboSpiralMatrixHandler()

	router.GET("/api/user/spiral", fiboSpiralMatrixHandler.GetSpiralMatrix)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/api/user/spiral", bytes.NewBuffer(jsonRequest))
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	req.URL.RawQuery = url.Values{
		"rows": {"4"},
		"cols": {"5"},
	}.Encode()

	c, _ := gin.CreateTestContext(w)

	c.Request = req

	var serviceAuthService = services.AuthServiceImpl{}

	err = serviceAuthService.HashUserPassword(c.Request.Context(), &signupRequestDto)
	assert.NoError(t, err)

	userTest := user.User{}
	userTest.FromModel(signupRequestDto)

	result := database.GlobalDB.Create(&userTest)
	assert.NoError(t, result.Error)

	router.ServeHTTP(w, req)

	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, 89, response.Rows[3][0])

	database.GlobalDB.Unscoped().Where("email = ?", signupRequestDto.Email).Delete(&user.User{})
}
