package services

import (
	"CodeBox/api/middleware"
	"CodeBox/models"
	"CodeBox/repository/db"
	"CodeBox/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func CreateAdmin(c *gin.Context) (interface{}, error) {
	var admin models.Admin
	err := c.BindJSON(&admin)
	if err != nil {
		return nil, err
	}
	admin.AdminId = uuid.New().String()
	if admin.Password == "" {
		return nil, fmt.Errorf("password can't be empty")
	}
	hashedPassword, err := utils.HashPassword(admin.Password)
	if err != nil {
		return nil, err
	}
	admin.Password = hashedPassword
	if err := validator.New().Struct(admin); err != nil {
		return nil, err
	}
	res, err := db.CreateOne(&admin)
	return res, err
}

func AdminSignIn(c *gin.Context) (interface{}, error) {
	var admin models.Admin
	err := c.BindJSON(&admin)
	if err != nil {
		return nil, fmt.Errorf("input miss match")
	}
	if admin.Username == "" || admin.Password == "" {
		return nil, fmt.Errorf("email or password can't be empty")
	}

	userPassword := admin.Password
	admin.Password = ""
	realUser, err := db.GetOne(admin)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err, fmt.Errorf("user not found")
		}

		return nil, err
	}

	err = utils.CheckPassword(userPassword, realUser.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "email or password is incorrect")
		return nil, fmt.Errorf("email or password is incorrect")
	}

	scope := []string{"admin"}
	token, err := middleware.GenerateToken(realUser.Username, "admin", scope)
	return token, err
}
