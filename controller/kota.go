package controller

import (
	"ilmi_backend/models"
	"ilmi_backend/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func (s *Server) GetAllKota(c *gin.Context) {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	kota := models.Kota{}

	idProvinsi := c.Query("provinsi")

	getAllKota, _ := kota.GetAllKotaTo(s.DB, idProvinsi)

	responseMessage := "Succes"

	response.GenericJsonResponse(c, http.StatusOK,
		responseMessage,
		getAllKota)

}
