package repository

import (
	"github.com/mariadb_golang/db"
	"github.com/mariadb_golang/helpers"
	"github.com/mariadb_golang/models"
)

func SaveEmployee(data models.Employee) error {
	db, err := db.InitDB()
	helpers.CheckError(err)
	defer db.Close()
	query := `INSERT INTO employees(branch_id, first_name, last_name, phone, salary)VALUES(?,?,?,?)`

	_, err = db.Exec(query, data.BranchID, data.FirstName, data.LastName, data.Phone, data.Salary)
	helpers.CheckError(err)
	return nil
}

func UpdateEmployee(data models.Employee  , id int) error {
	db, err := db.InitDB()
	helpers.CheckError(err)
	defer db.Close()
	query := `UPDATE employees SET branch_id=?, first_name=?, last_name=?, phone=?, salary=? WHERE id=?`
	_, err = db.Exec(query, data.BranchID,data.FirstName, data.LastName, data.Phone, data.Salary, id)
	helpers.CheckError(err)
	return nil
}

func GetAllEmployees()([]models.Employee, error){
	db, err := db.InitDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	query := `SELECT * FROM employees`
	

	rows, err := db.Query(query)
	defer rows.Close()
	var Employees []models.Employee
	for rows.Next(){
		var employee models.Employee
		err = rows.Scan(&employee.ID,&employee.BranchID, &employee.FirstName, &employee.LastName, &employee.Phone, &employee.Salary)
		if err != nil {
			return nil, err
		}
		Employees=append(Employees, employee)
	}
	return Employees, nil


}

func GetEmployeeById(id int) (models.Employee, error) {
	db, err := db.InitDB()
	if err != nil {
		return models.Employee{}, err
	}
	defer db.Close()
	query := `SELECT * FROM employees WHERE id =?`
	row := db.QueryRow(query, id)
	var employee models.Employee
	err = row.Scan(&employee.ID,&employee.BranchID, &employee.FirstName, &employee.LastName, &employee.Phone, &employee.Salary)
	if err != nil {
		return models.Employee{}, err
	}
	return employee, nil
}

func DeleteEmployee(id int) error{
	db, err := db.InitDB()
	helpers.CheckError(err)
	defer db.Close()
	query := `DELETE FROM employee WHERE id= ?`
	_, err = db.Exec(query, id)
	helpers.CheckError(err)
	return nil

}
