package controllers

import (
	"net/http"
	"rbac/models"
	"rbac/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserRoleController struct{}

func (UserRoleController) List(c *gin.Context) {
	user_roles := []models.UserRole{}
	utils.MysqlDB.Find(&user_roles)
	c.JSON(http.StatusOK, gin.H{
		"user_roles": user_roles,
	})
}

func (UserRoleController) Create(c *gin.Context) {
	user_id, _ := strconv.Atoi(c.PostForm("user_id"))
	role_id, _ := strconv.Atoi(c.PostForm("role_id"))

	user_role := models.UserRole{
		UserId: user_id,
		RoleId: role_id,
	}
	err := utils.MysqlDB.Create(&user_role).Error
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

func (UserRoleController) Delete(c *gin.Context) {
	user_id, _ := strconv.Atoi(c.PostForm("user_id"))
	role_id, _ := strconv.Atoi(c.PostForm("role_id"))

	// user_role := models.UserRole{
	// 	UserId: user_id,
	// 	RoleId: role_id,
	// }

	// utils.MysqlDB.Delete(&user_role)

	utils.MysqlDB.Where("user_id = ? AND role_id = ?", user_id, role_id).Delete(&models.UserRole{})

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
