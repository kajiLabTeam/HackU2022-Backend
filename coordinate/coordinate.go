package coordinate

import (
	"encoding/json"
	"example/createsql"
	"example/model"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/teris-io/shortid"
)

//coordinateのとき
func Coordinates(w http.ResponseWriter, r *http.Request) {
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
		// リクエストボディを読み込む
		/*
		   body, err := ioutil.ReadAll(r.Body)
		   if err != nil {
		       fmt.Fprintf(w, "Error: %v", err)
		       return
		   }
		   //構造体を定義
		   ble := model.Ble{}
		   // jsonを構造体に変換
		   err = json.Unmarshal(body, &ble)
		   if err != nil {
		       fmt.Fprintf(w, "Error: %v", err)
		       return
		   }
		*/

		w.Header().Set("Content-Type", "application/json")
		result := []*model.Coordinates{}
		db.Model(model.Coordinates{}).Where("ble = ? AND put_flag = ?", r.URL.Query().Get("ble"), 2).Find(&result)
		/*if err != nil {
		      json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
		      fmt.Fprintln(w, json_str)
		      return
		  }
		*/
		ble := model.Ble{}
		ble.Coordinate_id = result[0].Coordinate_id
		ble.Image = result[0].Image
		for i := 0; i < len(result); i++ {
			ble.Item[i].Category = result[i].Category
			ble.Item[i].Brand = result[i].Brand
			ble.Item[i].Price = result[i].Price
		}
		/*
		   for _, coordinate := range result {
		       js, err := json.Marshal(coordinate)
		       if err != nil {
		           //http.Error(w, err.Error(), http.StatusInternalServerError)
		           json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
		           fmt.Fprintln(w, json_str)
		           return
		       }
		       w.Write(js)
		       fmt.Println(coordinate)
		   }
		*/
		// SELECT * FROM coordinates WHERE ble = c1;

		//fmt.Println(result)
		var result1 model.Users
		err = db.Model(model.Users{}).Where("id = ?", result[0].User_id).First(&result1).Error
		if err != nil {
			json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
			fmt.Fprintln(w, json_str)
			return
		}
		ble.Users.Id = result1.Id
		ble.Users.Name = result1.Name
		ble.Users.Gender = result1.Gender
		ble.Users.Age = result1.Age
		ble.Users.Height = result1.Height
		ble.Users.Icon = result1.Icon
		js, err := json.Marshal(ble)
		if err != nil {
			//http.Error(w, err.Error(), http.StatusInternalServerError)
			json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
			fmt.Fprintln(w, json_str)
			return
		}

		w.Write(js)
		json_str := `{"status":"true"}`
		fmt.Fprintln(w, json_str)

	}

	//POSTのとき
	if r.Method == "POST" {

		// リクエストボディを読み込む
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
			fmt.Fprintln(w, json_str)
			return
		}
		//fmt.Fprintln(w, string(body))
		//構造体を定義
		clothes := model.CoordinatesAdd{}
		// jsonを構造体に変換
		err = json.Unmarshal(body, &clothes)
		if err != nil {
			json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
			fmt.Fprintln(w, json_str)
			return
		}
		//fmt.Fprintln(w, clothes)
		/*
			//構造体をjsonに変換
			json, err := json.Marshal(clothes)
			if err != nil {
				fmt.Fprintf(w, "Error: %v", err)
				return
			}
			fmt.Fprintln(w, string(json))
		*/

		//全てのput_flagを１にする
		createsql.UpdatePutFlag(db)
		//shortid作成
		sid, err := shortid.New(1, shortid.DefaultABC, 2342)
		if err != nil {
			json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
			fmt.Fprintln(w, json_str)
			return
		}
		//Coordinate_idを同じ服のとき同じにするため保存しておく
		shortId := sid.MustGenerate()
		uuid := uuid.New()
		for i := 0; i < len(clothes.Item); i++ {
			//それぞれのデータをとってきたデータにして登録
			err := db.Create(&model.Coordinates{
				Id:            sid.MustGenerate(),
				Coordinate_id: shortId,
				User_id:       clothes.User_id,
				Put_flag:      2,
				Public:        clothes.Public,
				Image:         clothes.Image,
				Category:      clothes.Item[i].Category,
				Brand:         clothes.Item[i].Brand,
				Price:         clothes.Item[i].Price,
				Ble:           uuid.String(),
				CreatedAt:     createsql.GetDate(),
				UpdatedAt:     createsql.GetDate(),
			}).Error
			if err != nil {
				json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
				fmt.Fprintln(w, json_str)
				return
			}
		}
		json_str := `{"status":"true"}`
		fmt.Fprintln(w, json_str)
		createsql.ShowCoordinate(db)
	}
}

func CoordinatesLike(w http.ResponseWriter, r *http.Request) {
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
		//fmt.Fprintln(w, string(body))
		//構造体を定義
		like := model.Likes{}
		// jsonを構造体に変換
		err = json.Unmarshal(body, &like)
		if err != nil {
			json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
			fmt.Fprintln(w, json_str)
			return
		}
		fmt.Fprintln(w, string(body))
		/*
			//構造体をjsonに変換
			json, err := json.Marshal(clothes)
			if err != nil {
				fmt.Fprintf(w, "Error: %v", err)
				return
			}
			fmt.Fprintln(w, string(json))
		*/
		//shortid作成
		sid, err := shortid.New(1, shortid.DefaultABC, 2342)
		if err != nil {
			json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
			fmt.Fprintln(w, json_str)
			return
		}
		//それぞれのデータをとってきたデータにして登録
		err = db.Create(&model.Likes{
			Id:            sid.MustGenerate(),
			Coordinate_id: like.Coordinate_id,
			Liked_user_id: like.Liked_user_id,
			User_id:       like.User_id,
			Lat:           like.Lat,
			Lng:           like.Lng,
			CreatedAt:     createsql.GetDate(),
			UpdatedAt:     createsql.GetDate(),
		}).Error
		if err != nil {
			json_str := `{"status":"false","message":"` + string(err.Error()) + `"}`
			fmt.Fprintln(w, json_str)
			return
		} else {
			json_str := `{"status":"true"}`
			fmt.Fprintf(w, json_str)

		}

		createsql.ShowLike(db)
	}
}
