package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func CheckError(text string, err interface{}) {
	if err != nil {
		fmt.Println(text, err)
	}
}

type User struct {
	Id       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Age      int    `json:"age,omitempty"`
	LastName string `json:"last_name,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Password string `json:"password,omitempty"`
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
	res, err := Get(db, 4)
	CheckError("error get: ", err)
	fmt.Println("Data inserted successfully.", res)
}

func Create(db *sql.DB, req *User) (*User, error) {
	var res = User{}
	query := `
INSERT INTO 
    users(first_name, age, last_name, phone, password) 
VALUES 
    (?, ?, ?, ?, ?)
RETURNING 
	id, first_name, last_name, phone
`
	err := db.QueryRow(query, req.Name, req.Age, req.LastName, req.Phone, req.Password).
		Scan(&res.Id, &res.Name, &res.LastName, &res.Phone)
	CheckError("error while inserting user info: ", err)
	return &res, err
}

func Get(db *sql.DB, id int) (*User, error) {
	query := `
SELECT id, first_name, last_name, age, phone FROM users WHERE id=?`
	res := User{}
	err := db.QueryRow(query, id).Scan(&res.Id, &res.Name, &res.LastName, &res.Age, &res.Phone)
	if err != nil {
		return &User{}, err
	}
	return &res, nil
}
