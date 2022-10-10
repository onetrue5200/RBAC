package controllers

import (
	"net/http"
	"rbac/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleController struct{}

func (RoleController) List(c *gin.Context) {
	roles := []models.Role{}
	models.MysqlDB.Where("is_delete=?", 0).Find(&roles)
	c.JSON(http.StatusOK, gin.H{
		"roles": roles,
	})
}

func (RoleController) Create(c *gin.Context) {
	name := c.PostForm("name")
	// check name
	roleList := []models.Role{}
	models.MysqlDB.Where("name=?", name).Find(&roleList)
	if len(roleList) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "name existed",
		})
		return
	}
	// create user
	role := models.Role{
		Name: name,
	}
	err := models.MysqlDB.Create(&role).Error
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
	models.MysqlDB.Where("name=?", name).Find(&roleList)
	if len(roleList) > 0 && roleList[0].Id != id {
		c.JSON(http.StatusOK, gin.H{
			"message": "name existed",
		})
		return
	}
	// update user
	role := models.Role{}
	models.MysqlDB.Where("ID=?", id).Find(&role)
	role.Name = name

	err := models.MysqlDB.Save(&role).Error
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
	id := c.PostForm("id")

	role := models.Role{}
	models.MysqlDB.Where("ID=?", id).Find(&role)

	role.IsDelete = 1

	err := models.MysqlDB.Save(&role).Error
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
