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
		json_str := &model.ErrorResponse{Status: false, Message: string(err.Error())}
		json, _ := json.Marshal(json_str)
		fmt.Fprintln(w, string(json))
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
		result := model.Users{}
		err = db.Where("mail = ?", r.URL.Query().Get("mail")).Find(&result).Error
		if err != nil {
			json_str := &model.ErrorResponse{Status: false, Message: string(err.Error())}
			json, _ := json.Marshal(json_str)
			fmt.Fprintln(w, string(json))
			return
		}
		result1 := model.UsersAddStatus{}
		result1.Id = result.Id
		result1.Name = result.Name
		result1.Gender = result.Gender
		result1.Age = result.Age
		result1.Height = result.Height
		result1.Uuid = result.Uuid
		result1.Mail = result.Mail
		result1.Icon = result.Icon

		result1.Status = true

		json, err := json.Marshal(result1)
		if err != nil {
			json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
			fmt.Fprintln(w, json_str)
			return
		}
		fmt.Fprintln(w, string(json))
	}

	//POSTのとき
	if r.Method == "POST" {
		// リクエストボディを読み込む
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			json_str := &model.ErrorResponse{Status: false, Message: string(err.Error())}
			json, _ := json.Marshal(json_str)
			fmt.Fprintln(w, string(json))
			return
		}
		//構造体を定義
		// var person Person
		user := model.Users{}
		// jsonを構造体に変換
		err = json.Unmarshal(body, &user)
		if err != nil {
			json_str := &model.ErrorResponse{Status: false, Message: string(err.Error())}
			json, _ := json.Marshal(json_str)
			fmt.Fprintln(w, string(json))
			return
		}
		sid, err := shortid.New(1, shortid.DefaultABC, 2342)
		if err != nil {
			json_str := &model.ErrorResponse{Status: false, Message: string(err.Error())}
			json, _ := json.Marshal(json_str)
			fmt.Fprintln(w, string(json))
			return
		}
		shortId := sid.MustGenerate()
		//それぞれのデータをとってきたデータにして登録
		err = db.Create(&model.Users{
			Id:        shortId,
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
			json_str := &model.ErrorResponse{Status: false, Message: string(err.Error())}
			json, _ := json.Marshal(json_str)
			fmt.Fprintln(w, string(json))
			return
			/*
				json_str := `{"status":false,"message":"` + string(error.Error()) + `"}`
				fmt.Fprintln(w, json_str)
			*/
		} else {
			json_str := &model.UserResponse{Status: true, Id: shortId}
			json, _ := json.Marshal(json_str)
			fmt.Fprintln(w, string(json))
		}
		createsql.ShowUser(db)

	}
}

func UsersMe(w http.ResponseWriter, r *http.Request) {
	db, err := createsql.SqlConnect()
	if err != nil {
		json_str := &model.ErrorResponse{Status: false, Message: string(err.Error())}
		json, _ := json.Marshal(json_str)
		fmt.Fprintln(w, string(json))
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
			json_str := &model.ErrorResponse{Status: false, Message: string(err.Error())}
			json, _ := json.Marshal(json_str)
			fmt.Fprintln(w, string(json))
			return
		}
		//構造体を定義
		// var user Users
		user := model.Users{}
		// jsonを構造体に変換
		err = json.Unmarshal(body, &user)
		if err != nil {
			json_str := &model.ErrorResponse{Status: false, Message: string(err.Error())}
			json, _ := json.Marshal(json_str)
			fmt.Fprintln(w, string(json))
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
			json_str := &model.ErrorResponse{Status: false, Message: string(err.Error())}
			json, _ := json.Marshal(json_str)
			fmt.Fprintln(w, string(json))
			return

		} else {
			json_str := &model.TrueResponse{Status: true}
			json, _ := json.Marshal(json_str)
			fmt.Fprintln(w, string(json))
		}
		createsql.ShowUser(db)
	}
}
