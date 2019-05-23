package main

import (
	"dto"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func gormConnect() *gorm.DB {
	DBMS := "postgres"
	USER := "appuser"
	PASS := "app1user"
	PROTOCOL := "192.168.99.100"
	DBNAME := "mstrndb"

	// CONNECT =
	db, err := gorm.Open(DBMS, "host="+PROTOCOL+" port=5432 sslmode=disable user="+USER+" password="+PASS+" dbname="+DBNAME)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	db := gormConnect()
	defer db.Close()

	userProfile := []dto.TUserProfile{}

	db.Find(&userProfile)

	fmt.Printf("取得したレコード数:%d\n", len(userProfile))

	for _, emp := range userProfile {
		fmt.Println("yaru")
		fmt.Println(emp.UserName)
		fmt.Println(emp.UserName)
		fmt.Println(emp.Birthday)
		fmt.Println(emp.MailAddress)
		fmt.Println(emp.Password)
		fmt.Println(emp.Profile)
		fmt.Println(emp.RegistData)
		fmt.Println(emp.CreateUser)
		fmt.Println(emp.CreateDate)
		fmt.Println(emp.UpdateUser)
		fmt.Println(emp.UpdateDate)
	}

}
