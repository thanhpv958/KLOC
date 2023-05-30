package controllers

import (
	"kloc-go/config/db"
	"kloc-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProjects(c *gin.Context) {
	var projects []models.Project

	name := c.Query("name")
	if name != "" {
		db.DB.Where("name LIKE ?", "%"+name+"%").Find(&projects)
	} else {
		db.DB.Find(&projects)
	}

	c.JSON(http.StatusOK, gin.H{"data": projects})
}
