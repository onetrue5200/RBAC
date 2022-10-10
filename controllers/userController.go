package controllers

import (
	"net/http"
	"rbac/models"
	"rbac/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (UserController) List(c *gin.Context) {
	users := []models.User{}
	models.MysqlDB.Where("is_delete=?", 0).Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (UserController) Create(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	// check username
	userList := []models.User{}
	models.MysqlDB.Where("username=?", username).Find(&userList)
	if len(userList) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "username existed",
		})
		return
	}
	// create user
	user := models.User{
		Username: username,
		Password: utils.Md5(password),
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

func (UserController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	username := c.PostForm("username")
	password := c.PostForm("password")
	// check if username existed
	userList := []models.User{}
	models.MysqlDB.Where("username=?", username).Find(&userList)
	if len(userList) > 0 && userList[0].Id != id {
		c.JSON(http.StatusOK, gin.H{
			"message": "username existed",
		})
		return
	}
	// update user
	user := models.User{}
	models.MysqlDB.Where("ID=?", id).Find(&user)
	user.Username = username
	user.Password = utils.Md5(password)

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

func (UserController) Delete(c *gin.Context) {
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
