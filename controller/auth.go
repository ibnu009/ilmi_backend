package controller

import (
	"errors"
	"fmt"
	"ilmi_backend/auth"
	"ilmi_backend/models"
	"ilmi_backend/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func (s *Server) Login(c *gin.Context) {
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

	token, err := s.SignIn(user.Email, user.Password)
	s.UpdateTokenUser(user.Email, token)

	fmt.Println(token)

	if err != nil {
		response.GenericJsonResponse(c, http.StatusBadRequest,
			err.Error(),
			nil)
		return
	}

	response.GenericJsonResponse(c, http.StatusOK,
		"Login Berhasil",
		models.TokenResponse{
			Token: token,
		})

}

// chek email & SigIn
func (s *Server) SignIn(email, password string) (string, error) {
	var err error
	user := models.User{}
	err = s.DB.Where("email = ?", email).Take(&user).Error
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
func (s *Server) UpdateTokenUser(email, token string) (string, error) {
	user := models.User{}
	err := s.DB.Model(&user).Where("email = ?", email).Update("token", token).Error
	if err != nil {
		return "", err
	}
	return token, nil
}
