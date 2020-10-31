package users

import (
	"github.com/apm-dev/go-user-rest-api/requests"
	"github.com/apm-dev/go-user-rest-api/services"
	"github.com/apm-dev/go-user-rest-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Index(c *gin.Context) {
	users, err := services.UserService.GetUsers()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users.Marshall())
}

func Show(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.BadRequest("user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	user, err := services.UserService.FindUser(userId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, user.Marshall())
}

func Store(c *gin.Context) {
	var userStoreRequest requests.UserStoreRequest
	if err := c.ShouldBindJSON(&userStoreRequest); err != nil {
		restErr := errors.ValidationError(userStoreRequest, err.Error())
		c.JSON(restErr.Status, restErr)
		return
	}
	user, saveErr := services.UserService.StoreUser(&userStoreRequest)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, user.Marshall())
}

func Update(c *gin.Context) {
	id, err := getUrlParamUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	var userUpdateRequest requests.UserUpdateRequest
	if err := c.ShouldBindJSON(&userUpdateRequest); err != nil {
		restErr := errors.ValidationError(userUpdateRequest, err.Error())
		c.JSON(restErr.Status, restErr)
		return
	}
	user, err := services.UserService.UpdateUser(
		&userUpdateRequest, id,
		c.Request.Method == http.MethodPatch,
	)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, user.Marshall())
}

func Delete(c *gin.Context) {
	id, err := getUrlParamUserID(c)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	err = services.UserService.DeleteUser(id)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": true})
}

func getUrlParamUserID(c *gin.Context) (int64, *errors.RestError) {
	id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return 0, errors.BadRequest("user id should be a number")
	}
	return id, nil
}
