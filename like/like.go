package like

import (
	"encoding/json"
	"example/createsql"
	"example/model"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Like(w http.ResponseWriter, r *http.Request) {
	db, err := createsql.SqlConnect()
	if err != nil {
		json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
		fmt.Fprintln(w, json_str)
		return
	} else {
		log.Print("seikou!")
	}
	defer db.Close()

	//GETのとき
	if r.Method == "GET" {

		//fmt.Println(r.URL.Query().Get("id"))
		/*
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
		*/

		//paramからidを受け取り、そのidのユーザーの服の情報、どこで評価されたかを全て表示
		//ユーザーの服の情報を返す
		w.Header().Set("Content-Type", "application/json")
		result := []*model.Coordinates{}
		err = db.Model(model.Coordinates{}).Where("user_id = ?", r.URL.Query().Get("id")).Find(&result).Error
		if err != nil {
			json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
			fmt.Fprintln(w, json_str)
			return
		}
		json, err := json.Marshal(result)
		if err != nil {
			json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
			fmt.Fprintln(w, json_str)
			return
		}

		fmt.Fprintln(w, string(json))
		json_str := `{"status":"true"}`
		fmt.Fprintln(w, json_str)

		/*
			for _, coordinate := range result {
				js, err := json.Marshal(coordinate)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(js)
				json_str = json_str
				fmt.Println(coordinate)
			}
		*/
		// SELECT * FROM coordinates WHERE ble = c1;

		//fmt.Println(result)

		/*
			//ユーザーがどこでどんな人に評価されたかを表示
			result1 := []*model.Likes{}
			db.Model(model.Likes{}).Where("user_id = ?", r.URL.Query().Get("id")).Find(&result1)
			for _, like := range result1 {
				js, err := json.Marshal(like)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(js)
				fmt.Println(like)
			}

			//fmt.Println(user.Id)

			//fmt.Println(result)

			//評価したユーザーの情報を返す
			var result2 model.Users
			for _, like := range result1 {
				db.Model(model.Users{}).Where("id = ?", like.Liked_user_id).First(&result2)
				js, err := json.Marshal(result2)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(js)
				fmt.Println(result2)
			}
		*/
		/*

			//paramからidを受け取り、そのidのユーザーの服の情報、どこで評価されたかを全て表示
			//ユーザーの服の情報を返す
			result3 := []*model.Coordinates{}
			db.Model(model.Coordinates{}).Where("public = ?", r.URL.Query().Get("public")).Find(&result3)
			for _, coordinate := range result3 {
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

			//ユーザーがどこでどんな人に評価されたかを表示
			result4 := []*model.Likes{}
			for _, coordinate := range result3 {
				db.Model(model.Likes{}).Where("user_id = ?", coordinate.User_id).Find(&result4)
				for _, like := range result4 {
					js, err := json.Marshal(like)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					w.Write(js)
					fmt.Println(like)
				}
			}

			//fmt.Println(user.Id)

			//fmt.Println(result)

			//評価したユーザーの情報を返す
			var result5 model.Users
			for _, like := range result4 {
				db.Model(model.Users{}).Where("id = ?", like.Liked_user_id).First(&result5)
				js, err := json.Marshal(result5)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(js)
				fmt.Println(result5)
			}
			fmt.Fprint(w, "}")
		*/

	}

}
