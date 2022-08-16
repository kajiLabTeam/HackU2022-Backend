package main

import (
	"example/coordinate"
	"example/createsql"
	"example/like"
	"example/user"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// db接続
	db, err := createsql.SqlConnect()
	if err != nil {
		panic(err.Error())
	} else {
		log.Print("seikou!")
	}
	defer db.Close()
	//postProfile(db)
	//postLikes(db)
	//postCoordinate(db)
	//getBle(db)
	//createsql.Delete(db)
	//既に作られているuserテーブルの確認
	fmt.Println("user")
	createsql.ShowUser(db)
	//coordinateテーブルの確認
	fmt.Println("coordinate")
	createsql.ShowCoordinate(db)
	//likeテーブルの確認
	fmt.Println("like")
	createsql.ShowLike(db)
	//coordinate
	http.HandleFunc("/coordinate", coordinate.Coordinates)
	http.HandleFunc("/coordinates/{coordinate_id}/like", coordinate.CoordinatesLike)
	http.HandleFunc("/login", user.Login)
	http.HandleFunc("/users/me", user.UsersMe)
	http.HandleFunc("/likes", like.Like)
	http.ListenAndServe(":8080", nil)
}
