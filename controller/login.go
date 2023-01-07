package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"ilmi_backend/auth"
	formaterror "ilmi_backend/format-error"
	"ilmi_backend/models"
	"ilmi_backend/response"
	"io/ioutil"
	"net/http"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err, "")
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err, "")
		return
	}

	//user.Prepare()
	err = user.Validate("login")
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err, "")
		return
	}

	token, err := server.SignIn(user.Email, user.Password)
	server.UpdateTokenUser(user.Email, token)
	fmt.Println(token)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		response.ERROR(w, http.StatusUnprocessableEntity, formattedError, formattedError.Error())
		return
	}

	userDetail, err := user.GetUserByEmail(server.DB, user.Email)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		response.ERROR(w, http.StatusUnprocessableEntity, formattedError, "")
		return
	}
	fmt.Println("token :" + token)
	response.JSONLOGIN(w, http.StatusOK, token, userDetail.Id, userDetail.Email)
}

// chek email & SigIn
func (server *Server) SignIn(email, password string) (string, error) {
	var err error
	user := models.User{}
	err = server.DB.Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}

	hashedPassword := models.HashPasswordToSha256(password)
	if hashedPassword != user.Password {
		return "", errors.New("Password Salah")
	}

	return auth.CreateToken(user.Id)
}

// update token
func (server *Server) UpdateTokenUser(email, token string) (string, error) {
	user := models.User{}
	err := server.DB.Model(&user).Where("email = ?", email).Update("token", token).Error
	if err != nil {
		return "", err
	}
	return token, nil
}
