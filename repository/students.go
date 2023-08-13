package repository

import (
	"github.com/mariadb_golang/db"
	"github.com/mariadb_golang/helpers"
	"github.com/mariadb_golang/models"
)

func SaveStudents(data models.Student) error {
	db, err := db.InitDB()
	helpers.CheckError(err)
	defer db.Close()
	query := `INSERT INTO students (first_name, last_name, phone, branch_id, course_id) VALUES(?,?,?,?,?)`
	_, err = db.Exec(query, data.FirstName, data.LastName, data.Phone, data.BranchID, data.CourseID)
	helpers.CheckError(err)
	return nil
}
func UpdateStudents(data models.Student, id int) error {
	db, err := db.InitDB()
	helpers.CheckError(err)
	defer db.Close()
	query := `UPDATE students SET first_name=?, last_name=?, phone=?, branch_id=?, course_id=? WHERE id=?`
	_, err = db.Exec(query, data.FirstName, data.LastName, data.Phone, data.BranchID, data.CourseID, id)
	helpers.CheckError(err)
	return nil
}

func GetAllStudents()([]models.Student, error){
	db, err := db.InitDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	query := `SELECT * FROM students`
	rows, err := db.Query(query)
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
func GetAllStudentCourse(id int) ([]models.Course, error){
	db, err := db.InitDB()
	if err != nil {
		return nil , err
	}
	defer db.Close()
	query := `SELECT c.* FROM course AS c JOIN students AS s ON c.id = s.course_id WHERE s.id= ?`
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
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

func GetStudentById(id int) (models.Student, error){
	db,err := db.InitDB()
	if err != nil {

		return models.Student{}, err
	}
	defer db.Close()
	query := `SELECT * FROM students WHERE id =?`
	row:= db.QueryRow(query, id)
	var student models.Student
	err = row.Scan(&student.ID, &student.FirstName, &student.LastName,&student.Phone, &student.BranchID, &student.CourseID )
	if err != nil {

		return models.Student{}, err
	}
	return student, nil
}

func DeleteStudent(id int) error {
	db, err := db.InitDB()
	helpers.CheckError(err)
	defer db.Close()
	query := `DELETE FROM students WHERE id =?`
	_, err = db.Exec(query,id)
	helpers.CheckError(err)
	return nil
}