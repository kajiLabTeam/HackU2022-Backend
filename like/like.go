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

		w.Header().Set("Content-Type", "application/json")
		result_test := model.Coordinates{}
		result := []*model.Coordinates{}
		if r.URL.Query().Get("user_id") != "" {

			err = db.Where("user_id = ?", r.URL.Query().Get("user_id")).Find(&result_test).Error
			if err != nil {
				json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
				fmt.Fprintln(w, json_str)
				return
			}

			err = db.Where("user_id = ?", r.URL.Query().Get("user_id")).Find(&result).Error
			if err != nil {
				json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
				fmt.Fprintln(w, json_str)
				return
			}
		} else if r.URL.Query().Get("public") != "" {
			err = db.Where("public = ?", r.URL.Query().Get("public")).Find(&result_test).Error
			if err != nil {
				json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
				fmt.Fprintln(w, json_str)
				return
			}
			err = db.Where("public = ?", r.URL.Query().Get("public")).Find(&result).Error
			if err != nil {
				json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
				fmt.Fprintln(w, json_str)
				return
			}
		}

		coordinate_list := []*model.Coordinates{}
		coordinate_list = append(coordinate_list, &model.Coordinates{Coordinate_id: result[0].Coordinate_id, User_id: result[0].User_id, Image: result[0].Image})
		//coordinate_id_list := []string{string(result[0].Coordinate_id)}

		for i := 1; i < len(result); i++ {
			if result[i].Coordinate_id != result[i-1].Coordinate_id {
				//coordinate_id_list = append(coordinate_id_list, result[i].Coordinate_id)
				coordinate_list = append(coordinate_list, &model.Coordinates{Coordinate_id: result[i].Coordinate_id, User_id: result[i].User_id, Image: result[i].Image})
			}
		}
		//coordinate_idのある分だけforを回し、いいねをした人の情報を取得し格納する
		liked_data := []*model.Map{}
		for i := 0; i < len(coordinate_list); i++ {

			/*
				//エラーがでた時に返すよう
				result1_test := model.Likes{}
				err = db.Where("coordinate_id = ?", coordinate_list[i].Coordinate_id).Find(&result1_test).Error
				if err != nil {
					json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
					fmt.Fprintln(w, json_str)
					return
				}
			*/

			//服をいいねした人のuser_idとすれ違った位置を取得
			result1 := []*model.Likes{}
			err = db.Where("coordinate_id = ?", coordinate_list[i].Coordinate_id).Find(&result1).Error
			if err != nil {
				json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
				fmt.Fprintln(w, json_str)
				return
			}

			//取得したユーザーの情報、評価した座標を保存するようの配列
			liked_user := []*model.Liked_user{}
			//いいねをした人の分だけforを回す
			for j := 0; j < len(result1); j++ {
				result2 := model.Users{}
				err = db.Where("id = ?", result1[j].Liked_user_id).First(&result2).Error
				if err != nil {
					json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
					fmt.Fprintln(w, json_str)
					return
				}

				//ユーザーの性別、年齢、身長、評価した座標を保存
				liked_user = append(liked_user, &model.Liked_user{Gender: result2.Gender, Age: result2.Age, Height: result2.Height, Lat: result1[j].Lat, Lng: result1[j].Lng})
			}
			liked_data = append(liked_data, &model.Map{Coordinate_id: coordinate_list[i].Coordinate_id, User_id: coordinate_list[i].User_id, Image: coordinate_list[i].Image, Liked_users: liked_user})

		}
		liked_data_array := model.Maps{Maps: liked_data, Status: "true"}
		json, err := json.Marshal(liked_data_array)
		if err != nil {
			json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
			fmt.Fprintln(w, json_str)
			return
		}
		fmt.Fprintln(w, string(json))
	}

}
