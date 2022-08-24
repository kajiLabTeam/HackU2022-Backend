package main

import (
	"xclothes/controller"
	"xclothes/database"

	"github.com/gin-gonic/gin"
)

func connectDatabase() {
	db := database.Connect()
	defer db.Close()
	// database.DeleteAll(db)
	database.Migrate(db)
	database.ShowUser(db)
	database.ShowCoordinate(db)
	database.ShowLike(db)
}

func main() {
	connectDatabase()
	engine := gin.Default()
	usersEngine := engine.Group("/users")
	{
		usersEngine.POST("", controller.CreateUsers)
		usersEngine.GET("", controller.FindUsers)
		usersEngine.GET("/:id", controller.FindUsersById)
		usersEngine.PUT("/:id", controller.UpdateUsersById)
		usersEngine.DELETE("/:id", controller.DeleteUsersById)
	}
	coordinatesEngine := engine.Group("/coordinates")
	{
		coordinatesEngine.POST("", controller.CreateCoordinates)
		coordinatesEngine.GET("", controller.FindCoordinates)
		coordinatesEngine.GET("/:id", controller.FindCoordinatesById)
		coordinatesEngine.PUT("/:id", controller.UpdateCoordinatesById)
		coordinatesEngine.DELETE("/:id", controller.DeleteCoordinatesById)
	}
	LikesEngine := engine.Group("/likes")
	{
		LikesEngine.POST("", controller.CreateLikes)
		LikesEngine.GET("", controller.FindLikes)
		LikesEngine.GET("/:id", controller.FindLikesById)
		LikesEngine.PUT("/:id", controller.UpdateLikesById)
		LikesEngine.DELETE("/:id", controller.DeleteLikesById)
	}
	engine.Run(":3000")
}
