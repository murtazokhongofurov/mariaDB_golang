package storage

import (
	"database/sql"
	"github.com/mariadb_golang/models"
)

type Storage struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (r *Storage) CreateUser(req *models.User) error {
	err := r.db.QueryRow()
}

func Create(db *sql.DB, req *models.User) (*models.User, error) {
	var res = models.User{}
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
	return &res, err
}

func Update(db *sql.DB, id int) error {
	var res = models.User{Name: "Otabek", LastName: "Toshmukhammedov", Age: 25, Phone: "1234567"}
	query := "UPDATE users SET first_name=?, last_name=?, age=?, phone=? WHERE id=?"
	_, err := db.Exec(query, res.Name, res.LastName, res.Age, res.Phone, id)
	return err
}
func Delete(db *sql.DB, Id int) error {
	query := "DELETE FROM users WHERE id=?"
	_, err := db.Exec(query, Id)
	return err

}

func Get(db *sql.DB, id int) (*models.User, error) {
	query := `
SELECT id, first_name, last_name, age, phone FROM users WHERE id=?`
	res := models.User{}
	err := db.QueryRow(query, id).Scan(&res.Id, &res.Name, &res.LastName, &res.Age, &res.Phone)
	if err != nil {
		return &models.User{}, err
	}
	return &res, nil
}

func GetList(db *sql.DB) ([]models.User, error) {
	res := []models.User{}
	query := `
SELECT 
    id, first_name, COALESCE(last_name, ''), age, COALESCE(phone, '') 
FROM users`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		temp := models.User{}
		err = rows.Scan(&temp.Id, &temp.Name, &temp.LastName, &temp.Age, &temp.Phone)
		if err != nil {
			return nil, err
		}
		res = append(res, temp)
	}
	return res, nil
}
