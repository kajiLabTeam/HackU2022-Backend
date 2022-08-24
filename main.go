package main

import (
	"time"
	"xclothes/controller"
	"xclothes/database"

	"github.com/gin-contrib/cors"
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

	// cors設定の適用
	engine.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{"*", "http://localhost:3000"},

		// アクセスを許可したいHTTPメソッド
		AllowMethods: []string{"*", "GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},

		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{"*", "Access-Control-Allow-Headers", "Content-Length", "Content-Type", "Authorization"},

		// cookieなどの認証情報を含めるか否か(通常デフォルトfalseなので合わせました)
		AllowCredentials: false,

		// プリフライトリクエストのキャッシュ時間
		MaxAge: 12 * time.Hour,
	}))

	usersEngine := engine.Group("/users")
	{
		usersEngine.POST("/signup", controller.CreateUsers)
		usersEngine.GET("", controller.FindUsers)
		usersEngine.GET("/:id", controller.FindUsersById)
		usersEngine.GET("/mail/:mail", controller.FindUsersByMail)
		usersEngine.PUT("/:id", controller.UpdateUsersById)
		usersEngine.DELETE("/:id", controller.DeleteUsersById)
	}
	coordinatesEngine := engine.Group("/coordinates")
	{
		coordinatesEngine.POST("", controller.CreateCoordinates)
		coordinatesEngine.POST("/:id/likes", controller.CreateLikeByCoordinateId)
		coordinatesEngine.GET("", controller.FindCoordinates)
		coordinatesEngine.GET("/:id", controller.FindCoordinatesById)
		coordinatesEngine.GET("/ble/:uuid", controller.FindCoordinatesByBle)
		coordinatesEngine.GET("/:id/likes", controller.FindLikesByCoordinateId)
		coordinatesEngine.GET("/public/likes", controller.FindLikesByCoordinatePublic)
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
