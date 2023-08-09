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
	//res, err := Get(db, 4)
	//_ = Delete(db, 1)
	res, err := GetList(db)
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

func Update(db *sql.DB, id int) error {
	var res = User{Name: "Otabek", LastName: "Toshmukhammedov", Age: 25, Phone: "1234567"}
	query := "UPDATE users SET first_name=?, last_name=?, age=?, phone=? WHERE id=?"
	_, err := db.Exec(query, res.Name, res.LastName, res.Age, res.Phone, id)
	CheckError("Update Error:", err)
	return err
}
func Delete(db *sql.DB, Id int) error {
	query := "DELETE FROM users WHERE id=?"
	_, err := db.Exec(query, Id)
	CheckError("Delete Error", err)
	return nil

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

func GetList(db *sql.DB) ([]User, error) {
	res := []User{}
	query := `
SELECT 
    id, first_name, COALESCE(last_name, ''), age, COALESCE(phone, '') 
FROM users`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		temp := User{}
		err = rows.Scan(&temp.Id, &temp.Name, &temp.LastName, &temp.Age, &temp.Phone)
		if err != nil {
			return nil, err
		}
		res = append(res, temp)
	}
	return res, nil
}
