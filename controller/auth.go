package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"ilmi_backend/auth"
	"ilmi_backend/models"
	"ilmi_backend/response"
	"net/http"
)

func (server *Server) Login(c *gin.Context) {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		response.GenericJsonResponse(c, http.StatusBadRequest,
			err.Error(),
			nil)
		return
	}

	err := user.Validate("login")
	if err != nil {
		response.GenericJsonResponse(c, http.StatusBadRequest,
			err.Error(),
			nil)
		return
	}

	token, err := server.SignIn(user.Email, user.Password)
	server.UpdateTokenUser(user.Email, token)
	fmt.Println(token)
	if err != nil {
		response.GenericJsonResponse(c, http.StatusBadRequest,
			"Login Berhasil",
			models.TokenResponse{
				Token: token,
			})
		return
	}
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
