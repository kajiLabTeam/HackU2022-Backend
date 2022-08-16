package main

import (
	"example/coordinate"
	"example/createsql"
	"example/like"
	"example/user"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/teris-io/shortid"
)

func main() {
	// db接続
	db, err := sqlConnect()
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
	fmt.Println("user")
	createsql.ShowUser(db)
	fmt.Println("coordinate")
	createsql.ShowCoordinate(db)
	fmt.Println("like")
	createsql.ShowLike(db)
	http.HandleFunc("/coordinate", coordinate.Coordinates)
	http.HandleFunc("/coordinates/{coordinate_id}/like", coordinate.CoordinatesLike)
	http.HandleFunc("/login", user.Login)
	http.HandleFunc("/users/me", user.UsersMe)
	http.HandleFunc("/likes", like.Like)
	http.ListenAndServe(":8080", nil)
}

//usersテーブルのデータを作成
func postProfile(db *gorm.DB) {
	sid, err := shortid.New(1, shortid.DefaultABC, 2342)
	if err != nil {
		panic(err.Error())
	}

	error := db.Create(&Users{
		Id:        sid.MustGenerate(),
		Name:      "ugyf",
		Gender:    1,
		Age:       "20〜25",
		Height:    170,
		Uuid:      "retcfyvg",
		Mail:      "ercty@gmail.com",
		Icon:      "esrtyv.jpg",
		CreatedAt: getDate(),
		UpdatedAt: getDate(),
	}).Error
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("データ追加成功")
	}
}

//usersテーブルのデータを作成
func postLikes(db *gorm.DB) {
	sid, err := shortid.New(1, shortid.DefaultABC, 2342)
	if err != nil {
		panic(err.Error())
	}

	error := db.Create(&Likes{
		Id:            sid.MustGenerate(),
		Coordinate_id: sid.MustGenerate(),
		Liked_user_id: sid.MustGenerate(),
		User_id:       sid.MustGenerate(),
		Lat:           "45.48742",
		Lng:           "63.72539",
		CreatedAt:     getDate(),
		UpdatedAt:     getDate(),
	}).Error
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("データ追加成功")
	}
}

/*
//coordinatesテーブルのデータを作成
func postCoordinate(db *gorm.DB) {
	updatePutFlag(db)
	sid, err := shortid.New(1, shortid.DefaultABC, 2342)
	if err != nil {
		panic(err.Error())
	}
	error := db.Create(&Coordinates{
		Id:            sid.MustGenerate(),
		Coordinate_id: sid.MustGenerate(),
		User_id:       "BCa1FIptM",
		Put_flag:      2,
		Public:        1,
		Image:         "asdf.jpeg",
		Category:      "pants",
		Brand:         "gu",
		Price:         "1000〜3000",
		Ble:           "abc1",
		CreatedAt:     getDate(),
		UpdatedAt:     getDate(),
	}).Error
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("データ追加成功")
	}
}
*/

//すれ違った人の情報を取得
func getBle(db *gorm.DB) {
	var result Coordinates
	db.Model(Coordinates{}).Where("ble = ? AND put_flag = ?", "abc1", 2).First(&result)
	// SELECT * FROM coordinates WHERE ble = c1;
	fmt.Println(result)
	var result1 Users
	db.Model(Users{}).Where("id = ?", result.User_id).First(&result1)
	fmt.Println(result1)
}

//データを更新
func updatePutFlag(db *gorm.DB) {
	db.Model(Coordinates{}).Updates(Coordinates{
		Put_flag:  1,
		UpdatedAt: getDate(),
	})
}

//データを見る
func showData(db *gorm.DB) {
	//usersテーブルの名前を全て表示
	result := []*Coordinates{}
	error := db.Find(&result).Error
	if error != nil || len(result) == 0 {
		return
	}
	count := 0
	for _, user := range result {
		count++
		fmt.Println(user)
		fmt.Println(count)
	}
}

//データを挿入した日時
func getDate() string {
	const layout = "2006-01-02 15:04:05"
	now := time.Now()
	return now.Format(layout)
}

// SQLConnect DB接続
func sqlConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(localhost:33060)"
	DBNAME := "clothesdb"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	return gorm.Open(DBMS, CONNECT)
}

// Users ユーザー情報のテーブル情報
type Users struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Gender    int    `json:"gender"`
	Age       string `json:"age"`
	Height    int    `json:"height"`
	Uuid      string `json:"uuid"`
	Mail      string `json:"mail"`
	Icon      string `json:"icon"`
	CreatedAt string `json:"created_at" sql:"not null;type:date"`
	UpdatedAt string `json:"update_at" sql:"not null;type:date"`
}

// Coordinates コーディネート情報のテーブル情報
type Likes struct {
	Id            string `json:"id"`
	Coordinate_id string `json:"coordinate_id"`
	Liked_user_id string `json:"liked_user_id"`
	User_id       string `json:"user_id"`
	Lat           string `json:"lat"`
	Lng           string `json:"lng"`
	CreatedAt     string `json:"created_at" sql:"not null;type:date"`
	UpdatedAt     string `json:"update_at" sql:"not null;type:date"`
}

// Coordinates コーディネート情報のテーブル情報
type Coordinates struct {
	Id            string `json:"id"`
	Coordinate_id string `json:"coordinate_id"`
	User_id       string `json:"user_id"`
	Put_flag      int    `json:"put_flag"`
	Public        int    `json:"public"`
	Image         string `json:"image"`
	Category      string `json:"category"`
	Brand         string `json:"brand"`
	Price         string `json:"price"`
	Ble           string `json:"ble"`
	CreatedAt     string `json:"created_at" sql:"not null;type:date"`
	UpdatedAt     string `json:"update_at" sql:"not null;type:date"`
}

// ユーザー情報
type Requestuser struct {
	Name   string `json:"name"`
	Gender int    `json:"gender"`
	Age    string `json:"age"`
	Height int    `json:"height"`
	Uuid   string `json:"uuid"`
	Mail   string `json:"mail"`
	Icon   string `json:"icon"`
}

// 服の情報
type Clothes struct {
	User_id  string `json:"user_id"`
	Image    string `json:"image"`
	Category string `json:"category"`
	Brand    string `json:"brand"`
	Price    string `json:"price"`
	Public   int    `json:"public"`
}

//jjjjjjjj
// hogehoge hugohugo
//coordinatesに変更
