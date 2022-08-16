package user

import (
	"encoding/json"
	"example/createsql"
	"example/model"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/teris-io/shortid"
)

func Login(w http.ResponseWriter, r *http.Request) {
	db, err := createsql.SqlConnect()
	if err != nil {
		json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
		fmt.Fprintln(w, json_str)
		return
	} else {
		log.Print("seikou!")
	}
	defer db.Close()

	//POSTのとき
	if r.Method == "POST" {
		// リクエストボディを読み込む
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
			fmt.Fprintln(w, json_str)
			return
		}
		//構造体を定義
		// var person Person
		user := model.Users{}
		// jsonを構造体に変換
		err = json.Unmarshal(body, &user)
		if err != nil {
			json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
			fmt.Fprintln(w, json_str)
			return
		}
		sid, err := shortid.New(1, shortid.DefaultABC, 2342)
		if err != nil {
			json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
			fmt.Fprintln(w, json_str)
			return
		}
		//それぞれのデータをとってきたデータにして登録
		err = db.Create(&model.Users{
			Id:        sid.MustGenerate(),
			Name:      user.Name,
			Gender:    user.Gender,
			Age:       user.Age,
			Height:    user.Height,
			Uuid:      user.Uuid,
			Mail:      user.Mail,
			Icon:      user.Icon,
			CreatedAt: createsql.GetDate(),
			UpdatedAt: createsql.GetDate(),
		}).Error
		if err != nil {
			json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
			fmt.Fprintln(w, json_str)
			return
			/*
				json_str := `{"status":false,"message":"` + string(error.Error()) + `"}`
				fmt.Fprintln(w, json_str)
			*/
		} else {
			json_str := `{"status":"true"}`
			fmt.Fprintln(w, json_str)
		}
		createsql.ShowUser(db)

	}
}

func UsersMe(w http.ResponseWriter, r *http.Request) {
	db, err := createsql.SqlConnect()
	if err != nil {
		json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
		fmt.Fprintln(w, json_str)
		return
	} else {
		log.Print("seikou!")
	}
	defer db.Close()

	//PATCHのとき
	if r.Method == "PATCH" {
		// リクエストボディを読み込む
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
			fmt.Fprintln(w, json_str)
			return
		}
		//構造体を定義
		// var user Users
		user := model.Users{}
		// jsonを構造体に変換
		err = json.Unmarshal(body, &user)
		if err != nil {
			json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
			fmt.Fprintln(w, json_str)
			return
		}
		err = db.Model(model.Users{}).Where("id = ?", user.Id).Updates(model.Users{
			Name:      user.Name,
			Icon:      user.Icon,
			Gender:    user.Gender,
			Age:       user.Age,
			Height:    user.Height,
			UpdatedAt: createsql.GetDate(),
		}).Error
		if err != nil {
			json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
			fmt.Fprintln(w, json_str)
			return

			//fmt.Fprintln(w, "status:false")
		} else {
			json_str := `{"status":"true"}`
			fmt.Fprintln(w, json_str)
		}
		createsql.ShowUser(db)
	}
}
