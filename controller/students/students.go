package students

import (
	"encoding/json"
	"net/http"

	."github.com/mariadb_golang/helpers"
	. "github.com/mariadb_golang/models"
	"github.com/mariadb_golang/repository"
)

func Save(w http.ResponseWriter, r *http.Request){
	var student Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "Invalid Data type", http.StatusBadRequest)
		return
	}
	err := repository.SaveStudents(student)
	RespondWithError(w, err.Error(), err)
	RestSuccess(w, "Successfully saved")
}

func Update(w http.ResponseWriter, r *http.Request){
	var student Student
	id, err := GetIDFromRoute(r)
	RespondWithError(w, err.Error(), err)
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "Invalid json error", http.StatusOK)
		return
	}
	err = repository.UpdateStudents(student, id)
	RespondWithError(w, err.Error(), err)
	RestSuccess(w, "Successfully Updated!!!")
}

func Delete(w http.ResponseWriter, r *http.Request){
	id, err := GetIDFromRoute(r)
	RespondWithError(w, err.Error(), err)
	err = repository.DeleteStudent(id)
	RespondWithError(w, err.Error(), err)
	RestSuccess(w, "Student has been successfully deleted")
}

func GetAllStudent(w http.ResponseWriter, r *http.Request){
	data, err := repository.GetAllStudents()
	RespondWithError(w, err.Error(), err)
	RespondWithJSON(w, data)
}

func GetById(w http.ResponseWriter, r *http.Request){
	id, err := GetIDFromRoute(r)
	RespondWithError(w, err.Error(), err)
	data, err := repository.GetStudentById(id)
	RespondWithError(w, err.Error(), err)
	RespondWithJSON(w, data)
}

func GetAllCourse(w http.ResponseWriter, r *http.Request){
	id, err := GetIDFromRoute(r)
	RespondWithError(w, err.Error(), err)
	data, err := repository.GetAllStudentCourse(id)
	RespondWithError(w, err.Error(), err)
	RespondWithJSON(w, data)
}