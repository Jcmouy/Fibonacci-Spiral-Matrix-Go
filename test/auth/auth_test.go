package auth

import (
	"bytes"
	"encoding/json"
	"fibonacci-spiral-matrix-go/internal/api/handler"
	"fibonacci-spiral-matrix-go/internal/config/database"
	"fibonacci-spiral-matrix-go/internal/core/domain/user"
	"fibonacci-spiral-matrix-go/internal/wired"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	actualResult := handler.SignupResponse{}

	userTest := user.User{
		Username: "Test User",
		Email:    "jwt@email.com",
		Password: "secret",
		Role:     "USER",
	}

	authHandler, _ := wired.InitializeAuthHandler()

	payload, err := json.Marshal(&userTest)
	assert.NoError(t, err)

	request, err := http.NewRequest("POST", "/api/public/signup", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = request

	err = database.InitDatabase()
	assert.NoError(t, err)

	database.GlobalDB.AutoMigrate(&user.User{})

	authHandler.SignUp(c)

	assert.Equal(t, 200, w.Code)

	err = json.Unmarshal(w.Body.Bytes(), &actualResult)
	assert.NoError(t, err)
	assert.Equal(t, "User registered successfully!", actualResult.Message)
}

func TestSignUpInvalidJSON(t *testing.T) {
	user := "test"

	authHandler := handler.AuthHandler{}

	payload, err := json.Marshal(&user)
	assert.NoError(t, err)

	request, err := http.NewRequest("POST", "/api/public/signup", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = request

	authHandler.SignUp(c)

	assert.Equal(t, 400, w.Code)
}

func TestLogin(t *testing.T) {
	userTest := handler.LoginPayload{
		Username: "Test User",
		Email:    "jwt@email.com",
		Password: "secret",
	}

	authHandler, _ := wired.InitializeAuthHandler()

	payload, err := json.Marshal(&userTest)
	assert.NoError(t, err)

	request, err := http.NewRequest("POST", "/api/public/login", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = request

	err = database.InitDatabase()
	assert.NoError(t, err)

	database.GlobalDB.AutoMigrate(&user.User{})

	authHandler.Login(c)

	assert.Equal(t, 200, w.Code)

}

func TestLoginInvalidJSON(t *testing.T) {
	user := "test"

	authHandler, _ := wired.InitializeAuthHandler()

	payload, err := json.Marshal(&user)
	assert.NoError(t, err)

	request, err := http.NewRequest("POST", "/api/public/login", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = request

	authHandler.Login(c)

	assert.Equal(t, 400, w.Code)
}

func TestLoginInvalidCredentials(t *testing.T) {
	userTest := handler.LoginPayload{
		Username: "Test User",
		Email:    "jwt@email.com",
		Password: "invalid",
	}

	authHandler, _ := wired.InitializeAuthHandler()

	payload, err := json.Marshal(&userTest)
	assert.NoError(t, err)

	request, err := http.NewRequest("POST", "/api/public/login", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = request

	err = database.InitDatabase()
	assert.NoError(t, err)

	database.GlobalDB.AutoMigrate(&user.User{})

	authHandler.Login(c)

	assert.Equal(t, 401, w.Code)

	database.GlobalDB.Unscoped().Where("email = ?", userTest.Email).Delete(&user.User{})
}
