package controller

import (
	"fmt"
	"ilmi_backend/models"
	"ilmi_backend/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

// Create
func (s *Server) CreateUser(c *gin.Context) {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	var user models.User

	if err := c.BindJSON(&user); err != nil {
		response.GenericJsonResponse(c, http.StatusBadRequest,
			err.Error(),
			nil)
		return
	}

	if user.IsUserExist(s.DB) {
		response.GenericJsonResponse(c, http.StatusBadRequest,
			"Akun dengan email ini telah terdaftar",
			nil)
		return
	}

	hashedPassword := models.HashPasswordToSha256(user.Password)
	user.Password = hashedPassword

	_, err := user.CreateUser(s.DB)
	if err != nil {
		response.GenericJsonResponse(c, http.StatusInternalServerError,
			err.Error(),
			nil)
		return
	}

	response.GenericJsonResponse(c, http.StatusCreated,
		"Berhasil Mendaftarkan Akun Ilmi",
		nil)
}

// get profile by id
func (s *Server) GetProfile(c *gin.Context) {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	user := models.User{}

	profile, err := user.GetUserProfile(s.DB, c)

	if err != nil {
		fmt.Println(err)
	}

	response.GenericJsonResponse(c, http.StatusOK,
		"Succes", profile)

}
