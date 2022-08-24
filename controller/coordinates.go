package controller

import (
	"net/http"

	"xclothes/database"
	"xclothes/model"

	"github.com/gin-gonic/gin"
)

func CreateCoordinates(c *gin.Context) {
	var coordinate model.Coordinate
	// Validation Check
	if err := c.BindJSON(&coordinate); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	// Connect database
	db := database.Connect()
	defer db.Close()
	// Create coordinate
	database.UpdatePutFlag(db, coordinate.UserID)
	coordinate.ID = database.GenerateId()
	for index := range coordinate.Wears {
		coordinate.Wears[index].ID = database.GenerateId()
	}
	// coordinate.Wears = coordinate.Wears.map(wear => )
	if err := db.Create(&coordinate).Error; err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	// Response
	c.JSON(http.StatusCreated, coordinate)
}

func FindCoordinates(c *gin.Context) {
	var coordinates []model.Coordinate
	// Connect database
	db := database.Connect()
	defer db.Close()
	// Find coordinates
	if err := db.Model(&model.Coordinate{}).Preload("Wears").Find(&coordinates).Error; err != nil {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	// Response
	c.JSON(http.StatusOK, coordinates)
}

func FindCoordinatesById(c *gin.Context) {
	// Get path pram ":id"
	id := c.Param("id")
	var coordinate model.Coordinate
	// Connect database
	db := database.Connect()
	defer db.Close()
	// Find coordinates
	if err := db.Model(&model.Coordinate{}).Preload("Wears").First(&coordinate, "id = ?", id).Error; err != nil {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	// Response
	c.JSON(http.StatusOK, coordinate)
}

func FindCoordinatesByBle(c *gin.Context) {
	// Get path pram ":id"
	ble := c.Param("uuid")
	var coordinate model.Coordinate
	// Connect database
	db := database.Connect()
	defer db.Close()
	// Find coordinates
	if err := db.Model(&model.Coordinate{}).Preload("Wears").Where("ble = ?", ble).First(&coordinate).Error; err != nil {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	// Response
	c.JSON(http.StatusOK, coordinate)
}

func UpdateCoordinatesById(c *gin.Context) {
	// Get path pram ":id"
	id := c.Param("id")
	var coordinate model.Coordinate
	// Validation Check
	if err := c.BindJSON(&coordinate); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	// Connect database
	db := database.Connect()
	defer db.Close()
	// Update coordinate
	coordinate.ID = id
	if err := db.Model(&model.Coordinate{}).Omit("Wears").Where("id = ?", id).Error; err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	// Response
	c.JSON(http.StatusCreated, coordinate)
}

func DeleteCoordinatesById(c *gin.Context) {
	// Get path pram ":id"
	id := c.Param("id")
	// Connect database
	db := database.Connect()
	defer db.Close()
	// Delete coordinate
	if err := db.Delete(&model.Coordinate{}, "id = ?", id).Error; err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, "OK")
}
