package helper

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func CheckUserRole(c *gin.Context, role string) (err error) {
	userRole := c.GetString("role")
	err = nil
	if userRole != role {
		err = errors.New("Unauthorized to access this resource")
		return err
	}

	return err
}
