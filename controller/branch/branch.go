package branch

import (
	"encoding/json"
	"net/http"

	"github.com/mariadb_golang/helpers"
	"github.com/mariadb_golang/models"
	"github.com/mariadb_golang/repository"
)

func GetAllBrachCourse(w http.ResponseWriter, r *http.Request){
	branches, err := repository.GetAllBranches()
	helpers.RespondWithError(w, err.Error(),err)
	helpers.RespondWithJSON(w, branches)
}

func GetBranchById(w http.ResponseWriter, r *http.Request){
	id, err := helpers.GetIDFromRoute(r)
	helpers.RespondWithError(w, err.Error(), err)
	data, err := repository.GetBranchById(id)
	helpers.RespondWithError(w, err.Error(), err)
	helpers.RespondWithJSON(w, data )
}

func SaveBranch(w http.ResponseWriter, r *http.Request){
	var branch models.Branch
	if err := json.NewDecoder(r.Body).Decode(&branch); err != nil {
		http.Error(w, "Invalid data structure", http.StatusBadRequest)
	}
	err := repository.CreateBranch(branch)
	helpers.RespondWithError(w, err.Error(), err)
	helpers.RestSuccess(w, "Successfully saved")

}
func UpdateBranch(w http.ResponseWriter, r *http.Request){
	id, err := helpers.GetIDFromRoute(r)
	helpers.RespondWithError(w, err.Error(), err)
	var branch models.Branch
	if err := json.NewDecoder(r.Body).Decode(&branch);  err != nil {
		http.Error(w, "Invalid data structure", http.StatusBadRequest)
	}
	err = repository.UpdateBranch(branch, id)
	helpers.RespondWithError(w, err.Error(), err)
	helpers.RestSuccess(w, "Branch successfully updated")
}

func DeleteBranch(w http.ResponseWriter, r *http.Request){
	id, err := helpers.GetIDFromRoute(r)
	helpers.RespondWithError(w, err.Error(), err)
	err = repository.DeleteBranch(id)
	helpers.RespondWithError(w, err.Error(), err)
	helpers.RestSuccess(w, "Successfully Deleted")
}
func GetAllBranchEmployees(w http.ResponseWriter, r *http.Request){
	id, err := helpers.GetIDFromRoute(r)
	helpers.RespondWithError(w, err.Error(), err)
	data , err := repository.GetAllBranchEmployees(id)
	helpers.RespondWithError(w, err.Error(), err)
	helpers.RespondWithJSON(w, data)
}

func GetAllBranchStudents(w http.ResponseWriter, r  *http.Request){
	id, err := helpers.GetIDFromRoute(r)
	helpers.RespondWithError(w, err.Error(), err)
	data , err := repository.GetAllBranchStudents(id)
	helpers.RespondWithError(w, err.Error(), err)
	helpers.RespondWithJSON(w, data)
}
func GetAllBranchCourses(w http.ResponseWriter, r  *http.Request){
	id, err := helpers.GetIDFromRoute(r)
	helpers.RespondWithError(w, err.Error(), err)
	data , err := repository.GetAllBrachCourse(id)
	helpers.RespondWithError(w, err.Error(), err)
	helpers.RespondWithJSON(w, data)
}
