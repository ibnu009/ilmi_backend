package controller

import (
	"encoding/json"
	"fmt"
	"ilmi_backend/models"
	"ilmi_backend/response"
	"io/ioutil"
	"net/http"
)

// Create
func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err, "")
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err, "")
		return
	}

	user.PrepareUser()
	hashedPassword := models.HashPasswordToSha256(user.Password)
	user.Password = hashedPassword

	userCreated, err := user.CreateUser(s.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err, "")
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.Id))
	response.JSON(w, http.StatusCreated, "Succes")
}
