package storage

import (
	"database/sql"
	"github.com/mariadb_golang/models"
)

func CreateBranch(db *sql.DB, req models.BranchReq) (models.BranchRes, error) {
	var res models.BranchRes
	query := `INSERT INTO branchs(name, addresses) VALUES(?, ?) RETURNING id, name, addresses`
	err := db.QueryRow(query, req.Name, req.Address).Scan(&res.Id, &res.Name, &res.Address)
	if err != nil {
		return models.BranchRes{}, err
	}
	return res, nil
}

func GetBranch(db *sql.DB, id int) (models.BranchRes, error) {
	var res models.BranchRes
	query := `SELECT id, name, addresses FROM branchs WHERE id=?`
	err := db.QueryRow(query, id).Scan(&res.Id, &res.Name, &res.Address)
	if err != nil {
		return models.BranchRes{}, err
	}
	return res, nil
}

func UpdateBranch(db *sql.DB, req models.BranchRes) error {
	query := `UPDATE branchs SET name=?, addresses=? WHERE id=?`
	_, err := db.Exec(query, req.Name, req.Address, req.Id)
	if err != nil {
		return err
	}
	return nil
}

func GetBranchList(db *sql.DB) ([]models.BranchRes, error) {
	var res []models.BranchRes
	query := `SELECT id, name, addresses FROM branchs`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		temp := models.BranchRes{}
		err = rows.Scan(&temp.Id, &temp.Name, &temp.Address)
		if err != nil {
			return nil, err
		}
		res = append(res, temp)
	}
	return res, nil

}

func GetBranchParams(db *sql.DB, key string) ([]models.BranchRes, error) {
	var res []models.BranchRes
	query := `SELECT id, name, addresses FROM branchs WHERE name LIKE ?`
	rows, err := db.Query(query, "%"+key+"%")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		temp := models.BranchRes{}
		err = rows.Scan(&temp.Id, &temp.Name, &temp.Address)
		if err != nil {
			return nil, err
		}
		res = append(res, temp)
	}
	return res, nil
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
