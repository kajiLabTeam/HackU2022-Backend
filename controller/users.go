package controller

import (
	"net/http"

	"xclothes/database"
	"xclothes/model"

	"github.com/gin-gonic/gin"
)

func CreateUsers(c *gin.Context) {
	var user model.User
	// Validation Check
	if err := c.BindJSON(&user); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	// Connect database
	db := database.Connect()
	defer db.Close()
	// Create coordinate
	user.ID = database.GenerateId()
	if err := db.Create(&user).Error; err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	// Response
	c.JSON(http.StatusCreated, user)
}

func FindUsers(c *gin.Context) {
	var users []model.User
	// Connect database
	db := database.Connect()
	defer db.Close()
	// Find coordinates
	if err := db.Find(&users).Error; err != nil {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	// Response
	c.JSON(http.StatusOK, users)
}

func FindUsersByMail(c *gin.Context) {
	// Get path pram ":mail"
	mail := c.Param("mail")
	var user model.User
	// Connect database
	db := database.Connect()
	defer db.Close()
	// Find coordinates
	if err := db.Where("mail = ?", mail).First(&user).Error; err != nil {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	// Response
	c.JSON(http.StatusOK, user)
}

func FindUsersById(c *gin.Context) {
	// Get path pram ":id"
	id := c.Param("id")
	var user model.User
	// Connect database
	db := database.Connect()
	defer db.Close()
	// Find coordinates
	if err := db.First(&user, "id = ?", id).Error; err != nil {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	// Response
	c.JSON(http.StatusOK, user)
}

func UpdateUsersById(c *gin.Context) {
	// Get path pram ":id"
	id := c.Param("id")
	var user model.User
	// Validation Check
	if err := c.BindJSON(&user); err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	// Connect database
	db := database.Connect()
	defer db.Close()
	// Update coordinate
	user.ID = id
	if err := db.Save(&user).Error; err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	// Response
	c.JSON(http.StatusCreated, user)
}

func DeleteUsersById(c *gin.Context) {
	// Get path pram ":id"
	id := c.Param("id")
	// Connect database
	db := database.Connect()
	defer db.Close()
	// Delete coordinate
	if err := db.Delete(&model.User{}, "id = ?", id).Error; err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, "OK")
}
