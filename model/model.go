package model

import (
	_ "github.com/go-sql-driver/mysql"
)

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

//ユーザーの情報をresponseで返すための構造体
type UsersAddStatus struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Gender    int    `json:"gender"`
	Age       string `json:"age"`
	Height    int    `json:"height"`
	Uuid      string `json:"uuid"`
	Mail      string `json:"mail"`
	Icon      string `json:"icon"`
	Status    string `json:"status"`
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

type CoordinatesAdd struct {
	Id            string `json:"id"`
	Coordinate_id string `json:"coordinate_id"`
	User_id       string `json:"user_id"`
	Put_flag      int    `json:"put_flag"`
	Public        int    `json:"public"`
	Image         string `json:"image"`
	Item          []Item `json:item`

	Ble string `json:"ble"`
}

type Item struct {
	Category string `json:"category"`
	Brand    string `json:"brand"`
	Price    string `json:"price"`
}
