package repository

import (
	"github.com/mariadb_golang/db"
	"github.com/mariadb_golang/helpers"
	"github.com/mariadb_golang/models"
)

func CRUD() {

}

func CreateBranch(data models.Branch) error {
	db , err := db.InitDB()
	helpers.CheckError(err)
	defer db.Close()
	query := "INSERT INTO branch (name, address) VALUES(?,?)"
	_,err = db.Exec(query, data.Addresses, data.Name)
	helpers.CheckError(err)
	return nil
}

func UpdateBranch( data models.Branch, Id int) error {
	db , err := db.InitDB()
	helpers.CheckError(err)
	defer db.Close()
	query := `UPDATE branch SET name=? , address=? WHERE id=?`
	_,err = db.Exec(query, data.Name, data.Name, Id) 
	helpers.CheckError(err)
	return nil
}

func GetAllBranches()([]models.Branch, error){
	db, err := db.InitDB()
	helpers.CheckError(err)
	defer db.Close()
	query := `SELECT * FROM branch`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var branches []models.Branch
	for rows.Next(){
		var branch models.Branch
		err = rows.Scan(&branch.ID, &branch.Name, &branch.Addresses)
		if err != nil {
			return nil , err
		}
		branches= append(branches, branch)
	}
	return branches, nil
}

func GetBranchById(id int) (models.Branch, error) {
	db, err := db.InitDB()
	if err != nil {
		return models.Branch{}, err
	}
	defer db.Close()
	query := `SELECT * FROM branch WHERE id=?`
	var branch models.Branch
	row := db.QueryRow(query, id)
	err = row.Scan(&branch.ID, &branch.Name, &branch.Addresses)
	if err != nil {
		return models.Branch{}, err
	}
	return branch, nil

}
func GetAllBranchEmployees(id int) ([]models.Employee, error){
	db, err := db.InitDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	query := `SELECT e.* FROM employees AS e JOIN branch AS b ON e.branch_id = b.id WHERE b.id =?`
	rows, err := db.Query(query, id)
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
func GetAllBranchStudents(id int) ([]models.Student, error){
	db, err := db.InitDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	query := `SELECT s.* FROM students AS s JOIN branch AS b ON s.branch_id = b.id WHERE b.id =?`
	rows, err := db.Query(query, id)
	defer rows.Close()
	var Students []models.Student
	for rows.Next(){
		var student models.Student
		err = rows.Scan(&student.ID, &student.FirstName, &student.LastName,&student.Phone, &student.BranchID, &student.CourseID )
		if err != nil {
			return nil , err
		}
		Students= append(Students, student)
	}
	return Students, nil
}

func GetAllBrachCourse(id int) ([]models.Course, error){
	db, err := db.InitDB()
	if err != nil {
		return nil , err
	}
	defer db.Close()
	query := `SELECT c.* FROM course AS c JOIN branch AS b ON c.branch_id = b.id WHERE b.id =?`
	rows, err := db.Query(query, id )
	if err != nil {
		return  nil , err
	}
	defer rows.Close()
	var Courses []models.Course
	for rows.Next(){
		var course models.Course
		err = rows.Scan(&course.ID, &course.Name, &course.BranchID)
		if err != nil {
			return nil, err
		}
		Courses = append(Courses, course)
	}
	return Courses, nil 
}

func DeleteBranch(id int) error{
	db, err := db.InitDB()
	helpers.CheckError(err)
	defer db.Close()
	query := `DELETE FROM branch WHERE id = ?`

	_,err = db.Exec(query, id)
	helpers.CheckError(err)
	return nil
}