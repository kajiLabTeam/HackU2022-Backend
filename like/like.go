package like

import (
	"encoding/json"
	"example/createsql"
	"example/model"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Like(w http.ResponseWriter, r *http.Request) {
	db, err := createsql.SqlConnect()
	if err != nil {
		panic(err.Error())
	} else {
		log.Print("seikou!")
	}
	defer db.Close()

	//GETのとき
	if r.Method == "GET" {

		fmt.Println(r.URL.Query())
		// リクエストボディを読み込む
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			return
		}
		//構造体を定義
		user := model.Users{}
		// jsonを構造体に変換
		err = json.Unmarshal(body, &user)
		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		result := []*model.Coordinates{}
		db.Model(model.Coordinates{}).Where("user_id = ?", user.Id).Find(&result)
		for _, coordinate := range result {
			js, err := json.Marshal(coordinate)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(js)
			fmt.Println(coordinate)
		}
		// SELECT * FROM coordinates WHERE ble = c1;

		//fmt.Println(result)

		result1 := []*model.Likes{}
		db.Model(model.Likes{}).Where("user_id = ?", user.Id).Find(&result1)
		for _, like := range result1 {
			js, err := json.Marshal(like)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(js)
			fmt.Println(like)
		}

		fmt.Println(user.Id)

		//fmt.Println(result)

		var result2 model.Users
		db.Model(model.Users{}).Where("id = ?", user.Id).First(&result2)
		js, err := json.Marshal(result2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(js)
		fmt.Println(result2)

	}

}
