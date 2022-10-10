package controllers

import (
	"net/http"
	"rbac/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (UserController) GetUsers(c *gin.Context) {
	users := []models.User{}
	models.MysqlDB.Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (UserController) CreateUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	userList := []models.User{}
	models.MysqlDB.Where("username=?", username).Find(&userList)
	if len(userList) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "username existed",
		})
		return
	}

	user := models.User{
		Username: username,
		Password: models.Md5(password),
	}
	err := models.MysqlDB.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "fail",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	}
}

func (UserController) UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	username := c.PostForm("username")
	password := c.PostForm("password")

	userList := []models.User{}
	models.MysqlDB.Where("username=?", username).Find(&userList)
	if len(userList) > 0 && userList[0].Id != id {
		c.JSON(http.StatusOK, gin.H{
			"message": "username existed",
		})
		return
	}

	user := models.User{}
	models.MysqlDB.Where("ID=?", id).Find(&user)
	user.Username = username
	user.Password = models.Md5(password)

	err := models.MysqlDB.Save(&user).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "fail",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	}
}

func (UserController) DeleteUser(c *gin.Context) {
	id := c.PostForm("id")

	user := models.User{}
	models.MysqlDB.Where("ID=?", id).Find(&user)

	user.IsDelete = 1

	err := models.MysqlDB.Save(&user).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "fail",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	}
}
