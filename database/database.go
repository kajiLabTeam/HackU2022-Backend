package database

import (
	"fmt"
	"time"

	"xclothes/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/teris-io/shortid"
)

func Connect() (database *gorm.DB) {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(localhost:33060)"
	DB_NAME := "x_clothes"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DB_NAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	if db, err := gorm.Open(DBMS, CONNECT); err != nil {
		panic(err.Error())
	} else {
		return db
	}
}

func GenerateId() string {
	if sid, err := shortid.New(1, shortid.DefaultABC, 2342); err != nil {
		panic(err.Error())
	} else {
		return sid.MustGenerate()
	}
}

func GetDateNow() time.Time {
	return time.Now()
}

func UpdatePutFlag(db *gorm.DB, user_id string) {
	db.Model(model.Coordinate{}).Where("user_id = ?", user_id).Updates(model.Coordinate{
		PutFlag:   false,
		UpdatedAt: GetDateNow(),
	})
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Coordinate{})
	db.AutoMigrate(&model.Wear{})
	db.AutoMigrate(&model.Like{})
}

func DeleteAll(db *gorm.DB) {
	db.Delete(&model.Coordinate{})
	db.Delete(&model.User{})
	db.Delete(&model.Like{})
}

func ShowUser(db *gorm.DB) {
	users := []*model.User{}
	if err := db.Find(&users).Error; err != nil {
		return
	}
	fmt.Println("=== users ===")
	for _, user := range users {
		fmt.Println(user)
	}
}

func ShowCoordinate(db *gorm.DB) {
	coordinates := []*model.Coordinate{}
	if err := db.Find(&coordinates).Error; err != nil {
		return
	}
	fmt.Println("=== coordinates ===")
	for _, coordinate := range coordinates {
		fmt.Println(coordinate)
	}
}

func ShowLike(db *gorm.DB) {
	Likes := []*model.Like{}
	if err := db.Find(&Likes).Error; err != nil {
		return
	}
	fmt.Println("=== likes ===")
	for _, like := range Likes {
		fmt.Println(like)
	}
}
