package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project-name-here/api/messages"
	"project-name-here/models"
	"project-name-here/usecase"
	"strconv"
)

var userService *usecase.UserService

func init() {

	userService = usecase.NewUserService()
}

func CreateUser(c *gin.Context) {

	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}

	created, err := userService.CreateUser(user)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": messages.ErrorsMessages[http.StatusUnprocessableEntity],
			"data":    struct{}{},
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": messages.SuccessMessages[http.StatusCreated],
		"data":    created,
	})
}

func ShowUser(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("user"))

	user, err := userService.ShowUser(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": messages.ErrorsMessages[http.StatusNotFound],
			"data":    struct{}{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": messages.SuccessMessages[http.StatusOK],
		"data":    user,
	})
}

func ListUsers(c *gin.Context) {

	users, err := userService.ListUsers()

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": messages.SuccessMessages[http.StatusUnprocessableEntity],
			"data":    users,
		})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusNoContent, gin.H{
			"message": messages.SuccessMessages[http.StatusOK],
			"data":    users,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": messages.SuccessMessages[http.StatusOK],
		"data":    users,
	})
}
