package controllers

import (
	"kloc-go/config/db"
	"kloc-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFoots(c *gin.Context) {
	var foots []models.Foot
	db.DB.Find(&foots)

	c.JSON(http.StatusOK, gin.H{"data": foots})
}
