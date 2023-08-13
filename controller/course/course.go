package course

import (
	"encoding/json"
	"net/http"

	."github.com/mariadb_golang/helpers"
	. "github.com/mariadb_golang/models"
	"github.com/mariadb_golang/repository"
)

func Save(w http.ResponseWriter, r *http.Request){
	var course Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}
	err := repository.SaveCourse(course)
	RespondWithError(w, err.Error(), err)
	RestSuccess(w, "Course has been successfully saved")
}

func Update(w http.ResponseWriter, r *http.Request){
	var course Course
	id, err := GetIDFromRoute(r)
	RespondWithError(w, err.Error(), err)
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}
	err = repository.UpdateCourse(course, id)
	RespondWithError(w, err.Error(), err)
	RestSuccess(w, "Course has been successfully updated")
}

func GetAll(w http.ResponseWriter, r *http.Request){
	data, err := repository.GetAllCourses()
	RespondWithError(w, err.Error(), err)
	RespondWithJSON(w, data)
}
func GetById(w http.ResponseWriter, r *http.Request){
	id, err := GetIDFromRoute(r)
	RespondWithError(w, err.Error(), err)
	data, err := repository.GetCourseById(id)
	RespondWithError(w, err.Error(), err)
	RespondWithJSON(w, data)
}

func Delete(w http.ResponseWriter, r *http.Request){
	id, err := GetIDFromRoute(r)
	RespondWithError(w, err.Error(), err)
	err = repository.DeleteCourse(id)
	RespondWithError(w, err.Error(), err)
	RestSuccess(w, "Successfully deleted!!")
}

func GetEnroledStudets(w http.ResponseWriter, r *http.Request){
	id, err := GetIDFromRoute(r)
	RespondWithError(w, err.Error(), err)
	data, err := repository.GetAllCourseEnrolledStudents(id)
	RespondWithError(w, err.Error(), err)
	RespondWithJSON(w, data)
}