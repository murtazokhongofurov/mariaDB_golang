package employee

import (
	"encoding/json"
	"net/http"

	. "github.com/mariadb_golang/helpers"
	. "github.com/mariadb_golang/models"
	"github.com/mariadb_golang/repository"
)

func Save(w http.ResponseWriter, r *http.Request){
	var employee Employee
	if err := json.NewDecoder(r.Body).Decode(employee); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}
	err := repository.SaveEmployee(employee)
	RespondWithError(w, err.Error(), err)
	RespondWithJSON(w, "Employee has successfully Saved")
}

func Update(w http.ResponseWriter, r *http.Request){
	var employee Employee
	id, err := GetIDFromRoute(r)
	RespondWithError(w, err.Error(), err)
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		http.Error(w, "Invalid json", http.StatusBadRequest)
	}
	err = repository.UpdateEmployee(employee, id)
	RespondWithError(w, err.Error(), err)
	RestSuccess(w, "Successfully Updated")
}

func GetAllEmployee(w http.ResponseWriter, r *http.Request){
	data, err := repository.GetAllEmployees()
	RespondWithError(w, err.Error(), err)
	RespondWithJSON(w, data)
}

func GetById(w http.ResponseWriter, r *http.Request){
	id, err := GetIDFromRoute(r)
	RespondWithError(w, err.Error(), err)
	data, err := repository.GetEmployeeById(id)
	RespondWithError(w, err.Error(), err)
	RespondWithJSON(w, data)
}

func Delete(w http.ResponseWriter, r *http.Request){
	id, err := GetIDFromRoute(r)
	RespondWithError(w, err.Error(), err)
	
}