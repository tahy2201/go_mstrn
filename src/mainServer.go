package main

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
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
