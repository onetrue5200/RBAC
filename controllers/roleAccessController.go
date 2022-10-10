package controllers

import (
	"net/http"
	"rbac/models"
	"rbac/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleAccessController struct{}

func (RoleAccessController) List(c *gin.Context) {
	data := []models.RoleAccess{}
	utils.MysqlDB.Find(&data)
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (RoleAccessController) Create(c *gin.Context) {
	role_id, _ := strconv.Atoi(c.PostForm("role_id"))
	access_id, _ := strconv.Atoi(c.PostForm("access_id"))

	data := models.RoleAccess{
		RoleId:   role_id,
		AccessId: access_id,
	}
	err := utils.MysqlDB.Create(&data).Error
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

func (RoleAccessController) Delete(c *gin.Context) {
	role_id, _ := strconv.Atoi(c.PostForm("role_id"))
	access_id, _ := strconv.Atoi(c.PostForm("access_id"))

	utils.MysqlDB.Where("role_id = ? AND access_id = ?", role_id, access_id).Delete(&models.RoleAccess{})

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
