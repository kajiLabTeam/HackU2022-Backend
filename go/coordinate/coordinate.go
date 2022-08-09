package coordinate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/teris-io/shortid"
)

//coordinateのとき
func Coordinates(w http.ResponseWriter, r *http.Request) {
	db, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	} else {
		log.Print("seikou!")
	}
	defer db.Close()

	//GETのとき
	if r.Method == "GET" {
		fmt.Fprintf(w, "Hello World! GET")
	}

	//POSTのとき
	if r.Method == "POST" {

		// リクエストボディを読み込む
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			return
		}
		//構造体を定義
		// var person Person
		clothes := Clothes{}
		// jsonを構造体に変換
		err = json.Unmarshal(body, &clothes)
		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			return
		}
		//構造体をjsonに変換
		/*
			json, err := json.Marshal(clothes)
			if err != nil {
				fmt.Fprintf(w, "Error: %v", err)
				return
			}*/

		//全てのput_flagを１にする
		updatePutFlag(db)
		//shortid作成
		sid, err := shortid.New(1, shortid.DefaultABC, 2342)
		if err != nil {
			panic(err.Error())
		}
		//Coordinate_idを同じ服のとき同じにするため保存しておく
		shortId := sid.MustGenerate()
		//それぞれのデータをとってきたデータにして登録
		error := db.Create(&Coordinates{
			Id:            sid.MustGenerate(),
			Coordinate_id: shortId,
			User_id:       clothes.User_id,
			Put_flag:      2,
			Public:        clothes.Public,
			Image:         clothes.Image,
			Category:      clothes.Category,
			Brand:         clothes.Brand,
			Price:         clothes.Price,
			CreatedAt:     getDate(),
			UpdatedAt:     getDate(),
		}).Error
		if error != nil {
			fmt.Println(error)
		} else {
			fmt.Println("データ追加成功")
		}
	}
}
