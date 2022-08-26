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
	//その入力されたuser_idのユーザが見つからなかった時、エラーを返す
	var user model.User
	if err := db.Model(&model.User{}).Where("id = ?", coordinate.UserID).First(&user).Error; err != nil {
		c.String(http.StatusBadRequest, "Bad request : Not Exist UserID")
		return
	}
	db.Model(&model.Coordinate{}).Where("user_id = ?", coordinate.UserID).Update("put_flag", false)
	coordinate.PutFlag = true
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
		c.String(http.StatusInternalServerError, "InternalServerError")
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

func FindCoordinatesByUserId(c *gin.Context) {
	// Get path pram ":id"
	id := c.Param("id")
	// Connect database
	db := database.Connect()
	defer db.Close()
	// Find coordinates
	var coordinate []model.Coordinate
	if err := db.Model(&model.Coordinate{}).Preload("Wears").Find(&coordinate, "user_id = ?", id).Error; err != nil {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	// Response
	c.JSON(http.StatusOK, coordinate)
}

func FindCoordinatesByBle(c *gin.Context) {
	// Get path pram ":id"
	ble := c.Param("uuid")
	var user model.User
	var coordinate model.Coordinate
	// Connect database
	db := database.Connect()
	defer db.Close()
	// Find coordinates
	if err := db.Model(&model.User{}).Where("ble = ?", ble).First(&user).Error; err != nil {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	if err := db.Model(&model.Coordinate{}).Preload("Wears").Where("user_id = ? AND put_flag = ?", user.ID, true).First(&coordinate).Error; err != nil {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	// Response
	c.JSON(http.StatusOK, coordinate)
}

//public=trueになっているcoordinateテーブルが全て帰ってくるもの
func FindCoordinatesByPublic(c *gin.Context) {
	var coordinates []model.Coordinate
	// Connect database
	db := database.Connect()
	defer db.Close()
	// Find coordinates
	if err := db.Model(&model.Coordinate{}).Preload("Wears").Where("public = ?", true).Find(&coordinates).Error; err != nil {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	// Response
	c.JSON(http.StatusOK, coordinates)
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
