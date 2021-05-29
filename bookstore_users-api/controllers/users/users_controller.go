package users

import (
	"net/http"
	"strconv"

	"github.com/HarryChang30/bookstore/bookstore_users-api/domain/users"
	"github.com/HarryChang30/bookstore/bookstore_users-api/services"
	"github.com/HarryChang30/bookstore/bookstore_utils-go/rest_errors"

	"github.com/gin-gonic/gin"
)

func getUserId(userIdParam string) (int64, rest_errors.RestErr) {
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		return 0, rest_errors.NewBadRequestError("user id should be a number")
	}

	return userId, nil
}

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

	c.JSON(http.StatusCreated, result.Marshall(true))
}

func GetUser(c *gin.Context) {}

func SearchUser(c *gin.Context) {
	status := c.Query("status")

	users, err := services.UsersService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(true))
}

func DeleteUser(c *gin.Context) {
	userId, err := getUserId(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	if err := services.UsersService.DeleteUser(userId); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
