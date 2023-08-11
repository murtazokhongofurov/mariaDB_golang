package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/mariadb_golang/models"
	"github.com/mariadb_golang/storage"
	"log"
	"net/http"
	"strconv"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong!!!"))
	}).Methods("GET")

	router.HandleFunc("/branch", CreateBranch).Methods("POST")
	router.HandleFunc("/branch/{id}", GetBranch).Methods("GET")
	router.HandleFunc("/branches", GetBranchList).Methods("GET")
	router.HandleFunc("/branches/{search}", GetBranchParams).Methods("GET")
	router.HandleFunc("/branch/update", UpdateBranch).Methods("PUT")

	if err := http.ListenAndServe(":4000", router); err != nil {
		log.Fatal("Error listen port: 4000")
	}

	dsn := "developer:password@tcp(localhost:3306)/testdb"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()
}

func CreateBranch(w http.ResponseWriter, r *http.Request) {
	dsn := "developer:password@tcp(localhost:3306)/testdb"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()

	var branch models.BranchReq
	err = json.NewDecoder(r.Body).Decode(&branch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := storage.CreateBranch(db, branch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(response)
}

func GetBranch(w http.ResponseWriter, r *http.Request) {
	dsn := "developer:password@tcp(localhost:3306)/testdb"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()

	vars := mux.Vars(r)["id"]
	id, err := strconv.Atoi(vars)
	if err != nil {
		log.Fatal("convert error: ", err)
	}
	res, err := storage.GetBranch(db, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(response)

}

func GetBranchList(w http.ResponseWriter, r *http.Request) {
	dsn := "developer:password@tcp(localhost:3306)/testdb"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()

	res, err := storage.GetBranchList(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(response)

}

func GetBranchParams(w http.ResponseWriter, r *http.Request) {
	dsn := "developer:password@tcp(localhost:3306)/testdb"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()
	key := mux.Vars(r)["search"]

	res, err := storage.GetBranchParams(db, key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(response)

}

func UpdateBranch(w http.ResponseWriter, r *http.Request) {
	dsn := "developer:password@tcp(localhost:3306)/testdb"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()

	var branch models.BranchRes
	err = json.NewDecoder(r.Body).Decode(&branch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = storage.UpdateBranch(db, branch)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var info = Message{
		Message: "Successfully updated",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response, err := json.Marshal(info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(response)
}

type Message struct {
	Message string `json:"message"`
}
