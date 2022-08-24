package model

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID        string    `json:"id"`
	UUID      string    `json:"uuid"`
	Mail      string    `json:"mail"`
	Name      string    `json:"name"`
	Gender    int       `json:"gender"`
	Age       string    `json:"age"`
	Height    int       `json:"height"`
	Icon      string    `json:"icon"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}

type Coordinate struct {
	ID        string    `json:"id"`
	PutFlag   bool      `json:"put_flag" gorm:"default:false"`
	Public    bool      `json:"public" gorm:"default:false"`
	Image     string    `json:"image"`
	Ble       string    `json:"ble"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
	Wears     []Wear    `json:"wears" gorm:"foreignKey:CoordinateID"`
}

// 服の情報
type Wear struct {
	ID           string    `json:"id"`
	Category     string    `json:"category"`
	Brand        string    `json:"brand"`
	Price        string    `json:"price"`
	CoordinateID string    `json:"coordinate_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"update_at"`
}

type Like struct {
	ID            string    `json:"id"`
	Lat           float32   `json:"lat"`
	Lon           float32   `json:"lng"`
	SendUserID    string    `json:"send_user_id"`
	ReceiveUserID string    `json:"receive_user_id"`
	CoordinateID  string    `json:"coordinate_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"update_at"`
}

// type Ble struct {
// 	Coordinate_id string  `json:"coordinate_id"`
// 	Image         string  `json:"image"`
// 	Items         []*Item `json:"items"`
// 	Users         Users   `json:"users"`
// 	Status        bool    `json:"status"`
// }

// type Send_user struct {
// 	Gender int    `json:"gender"`
// 	Age    string `json:"age"`
// 	Height int    `json:"height"`
// 	Lat    string `json:"lat"`
// 	Lng    string `json:"lng"`
// }

// type Map struct {
// 	Coordinate_id string       `json:"coordinate_id"`
// 	User_id       string       `json:"user_id"`
// 	Image         string       `json:"image"`
// 	Send_users    []*Send_user `json:"send_users"`
// }

// type Maps struct {
// 	Maps   []*Map `json:"map"`
// 	Status bool   `json:"status"`
// }

// type ErrorResponse struct {
// 	Status  bool   `json:"status"`
// 	Message string `json:"message"`
// }
// type UserResponse struct {
// 	Status bool   `json:"status"`
// 	Id     string `json:"id"`
// }
// type TrueResponse struct {
// 	Status bool `json:"status"`
// }
