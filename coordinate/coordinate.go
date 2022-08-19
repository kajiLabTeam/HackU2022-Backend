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

		w.Header().Set("Content-Type", "application/json")
		//fmt.Fprintln(w, r.URL.Query().Get("ble"))
		//クエリパラメータのbleUuidに一致していて、尚且つ今きている服の情報を取得
		result := []*model.Coordinates{}
		err := db.Model(model.Coordinates{}).Where("ble = ? AND put_flag = ?", r.URL.Query().Get("ble"), 2).Find(&result).Error
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

		//p := model.Item{Category: "category", Brand: "brand", Price: "price"}

		//服のid、服の写真、服のアイテム、ユーザー情報をまとめた構造体に変換し、json型にする
		ble := model.Ble{Coordinate_id: result[0].Coordinate_id, Image: result[0].Image, Items: p, Users: result1, Status: "true"}

		json, err := json.Marshal(ble)
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
		//fmt.Fprintln(w, clothes)
		for i := 0; i < len(clothes.Items); i++ {
			//それぞれのデータをとってきたデータにして登録
			err := db.Create(&model.Coordinates{
				Id:            sid.MustGenerate(),
				Coordinate_id: shortId,
				User_id:       clothes.User_id,
				Put_flag:      2,
				Public:        clothes.Public,
				Image:         clothes.Image,
				Category:      clothes.Items[i].Category,
				Brand:         clothes.Items[i].Brand,
				Price:         clothes.Items[i].Price,
				Ble:           clothes.Ble,
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
		//fmt.Fprintln(w, string(body))
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
			fmt.Fprintln(w, json_str)

		}

		createsql.ShowLike(db)
	}
}
