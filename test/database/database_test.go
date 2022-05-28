package database

import (
	"fibonacci-spiral-matrix-go/internal/config/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

//Cannot use 't' (type *testing.T) as the type TestingT - Bug in intellij

func TestInitDB(t *testing.T) {
	err := database.InitDatabase()
	assert.NoError(t, err)
}
