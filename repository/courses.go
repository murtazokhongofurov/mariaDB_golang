package repository

import (
	"github.com/mariadb_golang/db"
	"github.com/mariadb_golang/helpers"
	"github.com/mariadb_golang/models"
)

func SaveCourse(data models.Course) error{
	db, err := db.InitDB()
	helpers.CheckError(err)
	defer db.Close()
	query := `INSERT INTO course(name, branch_id) VALUES(?,?)`
	_, err = db.Exec(query, data.Name, data.BranchID)
	helpers.CheckError(err)
	return nil
}

func UpdateCourse(data models.Course, id int) error {
	db, err := db.InitDB()
	helpers.CheckError(err)
	defer db.Close()
	query := `UPDATE course SET name=?, branch_id=? WHERE id=?`
	_, err = db.Exec(query, data.Name, data.BranchID, id)
	helpers.CheckError(err)
	return nil
}
func DeleteCourse(id int) error {
	db, err := db.InitDB()
	helpers.CheckError(err)
	defer db.Close()
	query := `DELETE FROM course WHERE id =?`
	_, err = db.Exec(query, id)
	helpers.CheckError(err)
	return nil
}

func GetAllCourses()([]models.Course, error){
	db, err := db.InitDB()
	if err != nil {
		return nil , err
	}
	defer db.Close()
	query := `SELECT * from course`
	rows, err := db.Query(query)
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

func GetAllCourseEnrolledStudents(id int) ([]models.Student, error){
	db, err := db.InitDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	query := `SELECT s.* FROM students AS s JOIN course AS c ON c.id = s.course_id WHERE c.id= ?`
	rows, err := db.Query(query, id)
	if err != nil {
		return nil , err
	}
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
func GetCourseById(id int) (models.Course, error) {
	db, err := db.InitDB()
	if err != nil {
		return models.Course{} , err
	}
	defer db.Close()
	query := `SELECT * from course WHERE id =?`
	row := db.QueryRow(query, id)
	var course models.Course
	err = row.Scan(&course.ID, &course.Name, &course.BranchID)
	if err != nil {
		return models.Course{}, nil
	}
	return course, nil
}

