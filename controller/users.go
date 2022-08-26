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

//coordinateidをもとにその服を評価したlikesを受け取り、その中からsenduserを取り出しいいねをしたuser情報を返す
func FindUsersBySendUser(c *gin.Context) {
	// Get path pram ":id"
	id := c.Param("id")
	var likes []model.Like
	// Connect database
	db := database.Connect()
	defer db.Close()
	// Find Likes
	if err := db.Where("coordinate_id = ?", id).Find(&likes).Error; err != nil {
		c.String(http.StatusBadRequest, "Bad request : Not Exist CoordinateID")
		return
	}
	var sendUsers []string
	for _, senduser := range likes {
		sendUsers = append(sendUsers, senduser.SendUserID)
	}

	var users []model.User
	if err := db.Where("id IN (?)", sendUsers).Find(&users).Error; err != nil {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	// Response
	c.JSON(http.StatusOK, users)
}

func FindUsersByPublicSendUser(c *gin.Context) {
	var coordinates []model.Coordinate
	// Connect database
	db := database.Connect()
	defer db.Close()
	// Find coordinates
	if err := db.Where("public = ?", true).Find(&coordinates).Error; err != nil {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	var coordinateIds []string
	for _, coordinate := range coordinates {
		coordinateIds = append(coordinateIds, coordinate.ID)

	}

	var likes []model.Like

	if err := db.Where("coordinate_id IN (?)", coordinateIds).Find(&likes).Error; err != nil {
		c.String(http.StatusNotFound, "Not Found")
		return
	}

	var sendusers []string
	for _, like := range likes {
		sendusers = append(sendusers, like.SendUserID)

	}

	var users []model.User
	if err := db.Where("id IN (?)", sendusers).Find(&users).Error; err != nil {
		c.String(http.StatusNotFound, "Not Found")
		return
	}

	// Response
	c.JSON(http.StatusOK, users)
}
func FindUsersByReceiveUserIdLikesSendUser(c *gin.Context) {
	// Get path pram ":id"
	id := c.Param("id")
	var likes []model.Like
	// Connect database
	db := database.Connect()
	defer db.Close()
	// Find coordinates
	if err := db.Where("receive_user_id = ?", id).Find(&likes).Error; err != nil {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	var sendusers []string
	for _, like := range likes {
		sendusers = append(sendusers, like.SendUserID)

	}
	var users []model.User
	if err := db.Where("id IN (?)", sendusers).Find(&users).Error; err != nil {
		c.String(http.StatusNotFound, "Not Found")
		return
	}

	// Response
	c.JSON(http.StatusOK, users)
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
