package controllers

import (
	"net/http"
	"rbac/models"
	"rbac/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleController struct{}

func (RoleController) List(c *gin.Context) {
	roles := []models.Role{}
	utils.MysqlDB.Preload("Accesses").Where("is_delete=?", 0).Find(&roles)
	c.JSON(http.StatusOK, gin.H{
		"data": roles,
	})
}

func (RoleController) Create(c *gin.Context) {
	name := c.PostForm("name")
	// check name
	roleList := []models.Role{}
	utils.MysqlDB.Where("name=?", name).Find(&roleList)
	if len(roleList) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "角色名存在",
		})
		return
	}
	// create user
	role := models.Role{
		Name: name,
	}
	err := utils.MysqlDB.Create(&role).Error
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

func (RoleController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	name := c.PostForm("name")
	// check if username existed
	roleList := []models.Role{}
	utils.MysqlDB.Where("name=?", name).Find(&roleList)
	if len(roleList) > 0 && roleList[0].Id != id {
		c.JSON(http.StatusOK, gin.H{
			"message": "角色名存在",
		})
		return
	}
	// update user
	role := models.Role{}
	utils.MysqlDB.Where("ID=?", id).Find(&role)
	role.Name = name

	err := utils.MysqlDB.Save(&role).Error
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

func (RoleController) Delete(c *gin.Context) {
	id := c.Param("id")

	role := models.Role{}
	utils.MysqlDB.Where("ID=?", id).Find(&role)

	role.IsDelete = 1

	err := utils.MysqlDB.Save(&role).Error
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
