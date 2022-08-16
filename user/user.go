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
		panic(err.Error())
	} else {
		log.Print("seikou!")
	}
	defer db.Close()

	//GETのとき
	if r.Method == "POST" {
		// リクエストボディを読み込む
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			return
		}
		//構造体を定義
		// var person Person
		user := model.Users{}
		// jsonを構造体に変換
		err = json.Unmarshal(body, &user)
		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			return
		}
		sid, err := shortid.New(1, shortid.DefaultABC, 2342)
		if err != nil {
			panic(err.Error())
		}
		//それぞれのデータをとってきたデータにして登録
		error := db.Create(&model.Users{
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
		if error != nil {
			fmt.Println(error)
			fmt.Fprintf(w, "status:NG")
		} else {
			fmt.Fprintf(w, "status:OK")
			fmt.Println("データ追加成功")
		}
		createsql.ShowUser(db)

	}
}

func UsersMe(w http.ResponseWriter, r *http.Request) {
	db, err := createsql.SqlConnect()
	if err != nil {
		panic(err.Error())
	} else {
		log.Print("seikou!")
	}
	defer db.Close()

	//PATCHのとき
	if r.Method == "PATCH" {
		// リクエストボディを読み込む
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			return
		}
		//構造体を定義
		// var user Users
		user := model.Users{}
		// jsonを構造体に変換
		err = json.Unmarshal(body, &user)
		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			return
		}
		error := db.Model(model.Users{}).Where("id = ?", user.Id).Updates(model.Users{
			Name:      user.Name,
			Icon:      user.Icon,
			Gender:    user.Gender,
			Age:       user.Age,
			Height:    user.Height,
			UpdatedAt: createsql.GetDate(),
		}).Error
		if error != nil {
			fmt.Println(error)
			fmt.Fprintf(w, "status:NG")
		} else {
			fmt.Println("データ編集成功")
			fmt.Fprintf(w, "status:OK")
		}
		createsql.ShowUser(db)
	}
}
