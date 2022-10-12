package controllers

import (
	"net/http"
	"rbac/models"
	"rbac/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AccessController struct{}

func (AccessController) List(c *gin.Context) {
	data := []models.Access{}
	utils.MysqlDB.Where("is_delete=?", 0).Find(&data)
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (AccessController) Create(c *gin.Context) {
	name := c.PostForm("name")
	kind := c.PostForm("kind")
	url := c.PostForm("url")
	action := c.PostForm("action")
	module_id, _ := strconv.Atoi(c.PostForm("module_id"))

	dataList := []models.Access{}
	utils.MysqlDB.Where("name=?", name).Find(&dataList)
	if len(dataList) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "name existed",
		})
		return
	}

	data := models.Access{
		Name:     name,
		Kind:     kind,
		Url:      url,
		Action:   action,
		ModuleId: module_id,
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

func (AccessController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	name := c.PostForm("name")
	kind := c.PostForm("kind")
	url := c.PostForm("url")
	action := c.PostForm("action")
	module_id, _ := strconv.Atoi(c.PostForm("module_id"))

	dataList := []models.Access{}
	utils.MysqlDB.Where("name=?", name).Find(&dataList)
	if len(dataList) > 0 && dataList[0].Id != id {
		c.JSON(http.StatusOK, gin.H{
			"message": "name existed",
		})
		return
	}

	data := models.Access{}
	utils.MysqlDB.Where("ID=?", id).Find(&data)
	data.Name = name
	data.Kind = kind
	data.Url = url
	data.Action = action
	data.ModuleId = module_id

	err := utils.MysqlDB.Save(&data).Error
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

func (AccessController) Delete(c *gin.Context) {
	id := c.Param("id")

	data := models.Access{}
	utils.MysqlDB.Where("ID=?", id).Find(&data)

	data.IsDelete = 1

	err := utils.MysqlDB.Save(&data).Error
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
