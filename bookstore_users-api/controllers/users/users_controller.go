package users

import (
	"net/http"

	"github.com/HarryChang30/bookstore/bookstore_users-api/domain/users"
	"github.com/HarryChang30/bookstore/bookstore_users-api/services"
	"github.com/HarryChang30/bookstore/bookstore_utils-go/rest_errors"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	result, err := services.UsersService.CreateUser(user)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func GetUser(c *gin.Context) {}

func SearchUser(c *gin.Context) {}
