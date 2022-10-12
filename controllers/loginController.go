package controllers

import (
	"fmt"
	"net/http"
	"rbac/models"
	"rbac/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := models.User{}
	utils.MysqlDB.Where("username = ?", username).First(&user)
	fmt.Println(user.Password, utils.Md5(password))
	if utils.Md5(password) == user.Password {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "fail",
		})
	}
}
