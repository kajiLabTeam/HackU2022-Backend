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
func CreateLikeByCoordinateId(c *gin.Context) {
	var like model.Like
	// Validation Check
	id := c.Param("id")
	if err := c.BindJSON(&like); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	// Connect database
	db := database.Connect()
	defer db.Close()
	// Create coordinate
	like.ID = database.GenerateId()
	like.CoordinateID = id
	if err := db.Create(&like).Error; err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	// Response
	c.JSON(http.StatusCreated, like)
}

func FindLikes(c *gin.Context) {
	var likes []model.Like
	var filter model.Like
	// Get query pram "receive_user_id"
	if receiveUserId := c.Query("receive_user_id"); receiveUserId != "" {
		filter.ReceiveUserID = receiveUserId
	}
	// Connect database
	db := database.Connect()
	defer db.Close()
	// Find coordinates
	if err := db.Where(&filter).Find(&likes).Error; err != nil {
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

func FindLikesByCoordinateId(c *gin.Context) {
	// Get path pram ":id"
	id := c.Param("id")
	var likes []model.Like
	// Connect database
	db := database.Connect()
	defer db.Close()
	// Find coordinates
	if err := db.Where("coordinate_id = ?", id).Find(&likes).Error; err != nil {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	// Response
	c.JSON(http.StatusOK, likes)
}

func FindLikesByCoordinatePublic(c *gin.Context) {
	var coordinates []model.Coordinate
	// Connect database
	db := database.Connect()
	defer db.Close()
	// Find coordinates
	if err := db.Where("public = ?", true).Find(&coordinates).Error; err != nil {
		c.String(http.StatusNotFound, "Not Found 1")
		return
	}
	var coordinateIds []string
	for _, coordinate := range coordinates {
		coordinateIds = append(coordinateIds, coordinate.ID)

	}

	var likes []model.Like

	if err := db.Where("coordinate_id IN (?)", coordinateIds).Find(&likes).Error; err != nil {
		c.String(http.StatusNotFound, "Not Found 2")
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
