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
	utils.MysqlDB.Preload("Roles").Where("is_delete=?", 0).Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func (UserController) Create(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	// check username
	userList := []models.User{}
	utils.MysqlDB.Where("username=?", username).Find(&userList)
	if len(userList) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "用户名已存在",
		})
		return
	}
	// create user
	user := models.User{
		Username: username,
		Password: utils.Md5(password),
	}
	err := utils.MysqlDB.Create(&user).Error
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
	utils.MysqlDB.Where("username=?", username).Find(&userList)
	if len(userList) > 0 && userList[0].Id != id {
		c.JSON(http.StatusOK, gin.H{
			"message": "username existed",
		})
		return
	}
	// update user
	user := models.User{}
	utils.MysqlDB.Where("ID=?", id).Find(&user)
	user.Username = username
	if password == "" {
		user.Password = utils.Md5(password)
	}

	err := utils.MysqlDB.Save(&user).Error
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
	id := c.Param("id")

	user := models.User{}
	utils.MysqlDB.Where("ID=?", id).Find(&user)

	user.IsDelete = 1

	err := utils.MysqlDB.Save(&user).Error
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
