package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(map[string]interface{}{"status": statusCode, "data": data})
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func JSONLOGIN(w http.ResponseWriter, statusCode int, token interface{}, id interface{}, email interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(map[string]interface{}{"status": statusCode, "token": token, "id": id, "email": email})
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func JSONFILE(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(map[string]interface{}{"status": statusCode, "url": data})
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func JSONROWDATA(w http.ResponseWriter, statusCode int, data interface{}, count int64) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(map[string]interface{}{"status": statusCode, "data": data, "count": count})
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func JSONERROR(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error, message string) {
	if err != nil {
		JSONERROR(w, statusCode, map[string]interface{}{"status": statusCode, "error": err.Error(), "message": message})
		return
	}
	JSONERROR(w, http.StatusBadRequest, nil)
}

func ERRORQUERY(w http.ResponseWriter, statusCode int, err string) {
	if err != "" {
		JSONERROR(w, statusCode, map[string]interface{}{"status": statusCode, "error": err})
		return
	}
	JSONERROR(w, http.StatusBadRequest, nil)
}

func JSONSUCCESS(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(map[string]interface{}{"status": statusCode, "message": data})
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}
