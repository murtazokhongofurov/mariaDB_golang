package branch

import (
	"encoding/json"
	"net/http"

	"github.com/mariadb_golang/helpers"
	"github.com/mariadb_golang/models"
	"github.com/mariadb_golang/repository"
)

func GetAllBrach(w http.ResponseWriter, r *http.Request){
	branches, err := repository.GetAllBranches()
	if err != nil {

		helpers.RespondWithError(w, err.Error(),err)
		return
	}
	helpers.RespondWithJSON(w, branches)
}

func GetBranchById(w http.ResponseWriter, r *http.Request){
	id, err := helpers.GetIDFromRoute(r)
	if err!= nil{
		helpers.RespondWithError(w, err.Error(), err)
		return
	}
	data, err := repository.GetBranchById(id)
	if err != nil {
		helpers.RespondWithError(w, err.Error(), err)
		return
	}
	helpers.RespondWithJSON(w, data )
}

func SaveBranch(w http.ResponseWriter, r *http.Request){
	var branch models.Branch
	if err := json.NewDecoder(r.Body).Decode(&branch); err != nil {
		http.Error(w, "Invalid data structure", http.StatusBadRequest)
	}
	err := repository.CreateBranch(branch)
	if err != nil {

		helpers.RespondWithError(w, err.Error(), err)
		return
	}
	
	helpers.RestSuccess(w, "Successfully saved")

}
func UpdateBranch(w http.ResponseWriter, r *http.Request){
	id, err := helpers.GetIDFromRoute(r)
	if err != nil {

		helpers.RespondWithError(w, err.Error(), err)
		return
	}
	var branch models.Branch
	if err := json.NewDecoder(r.Body).Decode(&branch);  err != nil {
		http.Error(w, "Invalid data structure", http.StatusBadRequest)
	}
	err = repository.UpdateBranch(branch, id)
	if err != nil {

		helpers.RespondWithError(w, err.Error(), err)
		return
	}
	helpers.RestSuccess(w, "Branch successfully updated")
}

func DeleteBranch(w http.ResponseWriter, r *http.Request){
	id, err := helpers.GetIDFromRoute(r)
	if err != nil {

		helpers.RespondWithError(w, err.Error(), err)
		return
	}
	err = repository.DeleteBranch(id)
	if err != nil {

		helpers.RespondWithError(w, err.Error(), err)
		return
	}
	helpers.RestSuccess(w, "Successfully Deleted")
}
func GetAllBranchEmployees(w http.ResponseWriter, r *http.Request){
	id, err := helpers.GetIDFromRoute(r)
	if err != nil {

		helpers.RespondWithError(w, err.Error(), err)
		return
	}
	data , err := repository.GetAllBranchEmployees(id)
	if err != nil {

		helpers.RespondWithError(w, err.Error(), err)
		return
	}
	helpers.RespondWithJSON(w, data)
}

func GetAllBranchStudents(w http.ResponseWriter, r  *http.Request){
	id, err := helpers.GetIDFromRoute(r)
	if err != nil {

		helpers.RespondWithError(w, err.Error(), err)
		return
	}
	data , err := repository.GetAllBranchStudents(id)
	if err != nil {

		helpers.RespondWithError(w, err.Error(), err)
		return
	}
	helpers.RespondWithJSON(w, data)
}
func GetAllBranchCourses(w http.ResponseWriter, r  *http.Request){
	id, err := helpers.GetIDFromRoute(r)
	if err != nil {

		helpers.RespondWithError(w, err.Error(), err)
		return
	}
	data , err := repository.GetAllBrachCourse(id)
	if err != nil {

		helpers.RespondWithError(w, err.Error(), err)
		return
	}
	helpers.RespondWithJSON(w, data)
}
