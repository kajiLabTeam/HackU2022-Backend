package like

import (
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
		if r.URL.Query().Get("id") != "" {
			result_test := model.Coordinates{}
			err = db.Model(model.Coordinates{}).Where("user_id = ?", r.URL.Query().Get("id")).Find(&result_test).Error
			if err != nil {
				json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
				fmt.Fprintln(w, json_str)
				return
			}
			result := []*model.Coordinates{}
			err = db.Model(model.Coordinates{}).Where("user_id = ?", r.URL.Query().Get("id")).Find(&result).Error
			if err != nil {
				json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
				fmt.Fprintln(w, json_str)
				return
			}
			//服の情報にあるuser_idからユーザー情報を取得
			result1 := model.Users{}
			err = db.Model(model.Coordinates{}).Where("id = ?", result[0].User_id).First(&result1).Error
			if err != nil {
				json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
				fmt.Fprintln(w, json_str)
				return
			}

			//服の情報を登録しているデータベースから個別のアイテムのデータを配列に挿入
			p := []*model.Item{}
			for i := 0; i < len(result); i++ {
				p = append(p, &model.Item{Category: result[i].Category, Brand: result[i].Brand, Price: result[i].Price})
			}

			return
		}

		/*
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
		*/

	}

}
