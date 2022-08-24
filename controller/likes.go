package controller

import (
	"net/http"

	"xclothes/database"
	"xclothes/model"

	"github.com/gin-gonic/gin"
)

func CreateLikes(c *gin.Context) {
	var like model.Like
	// Validation Check
	if err := c.BindJSON(&like); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	// Connect database
	db := database.Connect()
	defer db.Close()
	// Create coordinate
	like.ID = database.GenerateId()
	if err := db.Create(&like).Error; err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	// Response
	c.JSON(http.StatusCreated, like)
}

func FindLikes(c *gin.Context) {
	var likes []model.Like
	// Connect database
	db := database.Connect()
	defer db.Close()
	// Find coordinates
	if err := db.Find(&likes).Error; err != nil {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	// Response
	c.JSON(http.StatusOK, likes)
}

func FindLikesById(c *gin.Context) {
	// Get path pram ":id"
	id := c.Param("id")
	var likes model.Like
	// Connect database
	db := database.Connect()
	defer db.Close()
	// Find coordinates
	if err := db.First(&likes, "id = ?", id).Error; err != nil {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	// Response
	c.JSON(http.StatusOK, likes)
}

func UpdateLikesById(c *gin.Context) {
	// Get path pram ":id"
	id := c.Param("id")
	var like model.Like
	// Validation Check
	if err := c.BindJSON(&like); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	// Connect database
	db := database.Connect()
	defer db.Close()
	// Update coordinate
	like.ID = id
	if err := db.Save(&like).Error; err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	// Response
	c.JSON(http.StatusCreated, like)
}

func DeleteLikesById(c *gin.Context) {
	// Get path pram ":id"
	id := c.Param("id")
	// Connect database
	db := database.Connect()
	defer db.Close()
	// Delete coordinate
	if err := db.Delete(&model.Like{}, "id = ?", id).Error; err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, "OK")
}
