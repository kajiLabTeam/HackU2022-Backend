package createsql

import (
	"example/model"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//データを更新
func UpdatePutFlag(db *gorm.DB, user_id string) {
	db.Model(model.Coordinates{}).Where("user_id = ?", user_id).Updates(model.Coordinates{
		Put_flag:  1,
		UpdatedAt: GetDate(),
	})
}

//データを見る
func ShowUser(db *gorm.DB) {
	//usersテーブルの名前を全て表示
	result := []*model.Users{}
	error := db.Find(&result).Error
	if error != nil || len(result) == 0 {
		return
	}
	for _, users := range result {
		fmt.Println(users)
	}
}

//データを見る
func ShowCoordinate(db *gorm.DB) {
	//usersテーブルの名前を全て表示
	result := []*model.Coordinates{}
	error := db.Find(&result).Error
	if error != nil || len(result) == 0 {
		return
	}
	for _, coordinates := range result {
		fmt.Println(coordinates)
	}
}

//データを見る
func ShowLike(db *gorm.DB) {
	//usersテーブルの名前を全て表示
	result := []*model.Likes{}
	error := db.Find(&result).Error
	if error != nil || len(result) == 0 {
		return
	}
	for _, like := range result {
		fmt.Println(like)
	}
}

//データを挿入した日時
func GetDate() string {
	const layout = "2006-01-02 15:04:05"
	now := time.Now()
	return now.Format(layout)
}

// SQLConnect DB接続
func SqlConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(localhost:33060)"
	DBNAME := "clothesdb"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	return gorm.Open(DBMS, CONNECT)
}

func Delete(db *gorm.DB) {
	error := db.Delete(model.Coordinates{}).Error
	//// DELETE from users where id=1
	if error != nil {
		fmt.Println(error)
	}
	db.Delete(model.Users{})
	db.Delete(model.Likes{})
}
