package helpers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ErrorResponse struct {
    Message string
}

func RestSuccess(w http.ResponseWriter, message string){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	errResponse := ErrorResponse{
		Message: message,
	}

	errJSON, _ := json.Marshal(errResponse)
	_, _ = w.Write(errJSON)
}

func RespondWithError(w http.ResponseWriter, message string, err error) {
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusInternalServerError)

        errResponse := ErrorResponse{
            Message: message,
        }

        errJSON, _ := json.Marshal(errResponse)
        _, _ = w.Write(errJSON)
        return
    }
}

func RespondWithJSON(w http.ResponseWriter, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(data); err != nil {
        RespondWithError(w, "Failed to encode response", err)
    }
}

func GetIDFromRoute(r *http.Request) (int, error) {
    vars := mux.Vars(r)
    idStr, ok := vars["id"]
    if !ok {
        return 0, http.ErrNoCookie
    }

    id, err := strconv.Atoi(idStr)
    if err != nil {
        return 0, err
    }
    
    return id, nil
}





