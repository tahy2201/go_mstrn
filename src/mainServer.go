package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/tahy2201/go_mstrn/src/dto"
)

// "net/http"

// type String string

// func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//   fmt.Fprint(w, s)
// }

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/get/userprof", GetUserProfile),
		rest.Post("/add/traindata", AddTrainData),
	// ここでは１つしかAPIがないが、例えば、
	// rest.POST("/api/comic", i.PostComicData),
	// みたいに書けば、処理を分岐させられる
	)
	// もし、ルーティングにエラーがあれば、ログを吐く
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

// GetUserProfile aaa
func GetUserProfile(w rest.ResponseWriter, r *rest.Request) {
	userProfs := ExecDb()
	w.WriteHeader(http.StatusOK)
	w.WriteJson(&userProfs)
}

// AddTrainData トレーニング記録追加
func AddTrainData(w rest.ResponseWriter, req *rest.Request) {

	input := dto.AddTrainData{}
	// 各種バリデーションが行われている README.md参照
	err := req.DecodeJsonPayload(&input)

	// バリデーションで引っかかったエラーを返す
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// if input.Name == "" {
	//     rest.Error(w, "Name is required", 400)
	// }

	nowDate := time.Now()

	var trainData dto.TTrainRecord
	trainData.ExerciseID = input.ExerciseID
	trainData.SetCount = input.SetCount
	trainData.UserID = input.UserID
	trainData.ExerciseWeight = input.ExerciseWeight
	trainData.ExerciseTimes = input.ExerciseTimes
	trainData.ExerciseMemo = input.ExerciseMemo
	trainData.CreateUser = input.UserID
	trainData.CreateDate = nowDate
	trainData.UpdateUser = input.UserID
	trainData.UpdateDate = nowDate

	InsertTrainRecord(&trainData)

	w.WriteJson("{\"status\":\"OK!!\"}\n")
}
