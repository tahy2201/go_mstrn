package main

import (
	"fmt"

	"github.com/tahy2201/go_mstrn/src/dto"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func gormConnect() *gorm.DB {
	DBMS := "postgres"
	USER := "appluser"
	PASS := "app1user"
	PROTOCOL := "localhost"
	DBNAME := "mstrndb"

	// CONNECT =
	db, err := gorm.Open(DBMS, "host="+PROTOCOL+" port=5432 sslmode=disable user="+USER+" password="+PASS+" dbname="+DBNAME)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func InsertTrainRecord(trainData *dto.TTrainRecord) {
	db := gormConnect()
	defer db.Close()

	trainData.TrainID = 0
	db.Create(&trainData)
}

// ExecDb DB実行結果を返します。
func ExecDb() *[]dto.TUserProfile {
	db := gormConnect()
	defer db.Close()

	userProfile := []dto.TUserProfile{}

	db.Debug().Find(&userProfile)

	fmt.Printf("取得したレコード数:%d\n", len(userProfile))

	fmt.Printf("取得結果%v\n", userProfile)

	// foreachっぽい書き方
	// for _, emp := range userProfile {
	// 	fmt.Println(emp.UserId)
	// }
	return &userProfile
}
