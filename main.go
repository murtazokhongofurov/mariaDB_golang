package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mariadb_golang/storage"
	"log"
)

func CheckError(text string, err interface{}) {
	if err != nil {
		fmt.Println(text, err)
	}
}

func main() {
	dsn := "developer:password@tcp(localhost:3306)/testdb"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()
	//user := User{
	//	Name:     "Komil",
	//	Age:      23,
	//	LastName: "Davlatov",
	//	Phone:    "+094983459",
	//	Password: "3434545",
	//}
	//res, err := Create(db, &user)
	//CheckError("error Create: ", err)
	//res, err := Get(db, 4)
	//_ = Delete(db, 1)
	res, err := storage.GetList(db)
	CheckError("error get: ", err)
	fmt.Println("Data inserted successfully.", res)
}
