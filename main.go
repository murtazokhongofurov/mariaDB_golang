package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mariadb_golang/controller/branch"
	"github.com/mariadb_golang/controller/course"
	"github.com/mariadb_golang/controller/employee"
	"github.com/mariadb_golang/controller/students"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/branch", branch.SaveBranch).Methods("POST")
	r.HandleFunc("/branch/{id}", branch.UpdateBranch).Methods("PUT")
	r.HandleFunc("/branch/{id}", branch.GetBranchById).Methods("GET")
	r.HandleFunc("/branches", branch.GetAllBrach).Methods("GET")
	r.HandleFunc("/branche-employees/{id}", branch.GetAllBranchEmployees).Methods("GET")
	r.HandleFunc("/branche-courses/{id}", branch.GetAllBranchCourses).Methods("GET")
	r.HandleFunc("/branche-students/{id}", branch.GetAllBranchStudents).Methods("GET")
	r.HandleFunc("/branch/{id}", branch.DeleteBranch).Methods("DELETE")

	r.HandleFunc("/employee", employee.Save).Methods("POST")
	r.HandleFunc("/employee/{id}", employee.Update).Methods("PUT")
	r.HandleFunc("/employee/{id}", employee.GetById).Methods("GET")
	r.HandleFunc("/employees", employee.GetAllEmployee).Methods("GET")
	r.HandleFunc("/employee-delete/{id}", employee.Delete).Methods("DELETE")

	r.HandleFunc("/course", course.Save).Methods("POST")
	r.HandleFunc("/course/{id}", course.Update).Methods("PUT")
	r.HandleFunc("/course/{id}", course.Delete).Methods("DELETE")
	r.HandleFunc("/course/{id}", course.GetById).Methods("GET")
	r.HandleFunc("/courses", course.GetAll).Methods("GET")
	r.HandleFunc("/course/{id}", course.GetById).Methods("GET")
	r.HandleFunc("/course-student/{id}", course.GetEnroledStudets).Methods("GET")

	r.HandleFunc("/student", students.Save).Methods("POST")
	r.HandleFunc("/student/{id}", students.Update).Methods("PUT")
	r.HandleFunc("/student/{id}", students.Delete).Methods("DELETE")
	r.HandleFunc("/students", students.GetAllStudent).Methods("GET")
	r.HandleFunc("/student/{id}", students.GetById).Methods("GET")
	r.HandleFunc("/student/{id}", students.GetAllCourse).Methods("GET")

	http.Handle("/", r)

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8000", nil)
}