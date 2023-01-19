package controller

import (
	"fmt"
	"ilmi_backend/models"
	"ilmi_backend/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func (s *Server) CreatedScheduleSholat(c *gin.Context) {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	ss := models.JadwalSholat{}

	if err := c.BindJSON(&ss); err != nil {
		response.GenericJsonResponse(c, http.StatusBadRequest,
			err.Error(),
			nil)
		return
	}
	responseMessage := "Succes"

	createdJadwalSholat, err := ss.CreateScheduleSholat(s.DB)
	if err != nil {
		responseMessage = "error Bad Request"
	}
	response.GenericJsonResponse(c, http.StatusOK,
		responseMessage,
		createdJadwalSholat)
}

func (s *Server) GetScheduleSholat(c *gin.Context) {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	jsSholat := models.JadwalSholat{}

	idProvinsi := c.Query("provinsi")
	idKota := c.Query("kota")

	js, err := jsSholat.GetScheduleSholatProvinsiKota(s.DB, idProvinsi, idKota)

	responseMessage := "Succes"
	if err != nil {
		responseMessage = "Tolong Pilih Provinsi Terlebih Dahulu"
	}

	response.GenericJsonResponse(c, http.StatusOK,
		responseMessage,
		js)
}

func (s *Server) UpdateScheduleSholat(c *gin.Context) {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	jsSholat := models.JadwalSholat{}

	if err := c.BindJSON(&jsSholat); err != nil {
		response.GenericJsonResponse(c, http.StatusBadRequest,
			err.Error(),
			nil)
		return
	}

	updateJs, err := jsSholat.UpdateScheduleSholat(s.DB, c)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
	}

	response.GenericJsonResponse(c, http.StatusOK,
		"Succes",
		updateJs)

}

func (s *Server) DeleteScheduleSholat(c *gin.Context) {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	jsSholat := models.JadwalSholat{}

	deleteJs, err := jsSholat.DeleteScheduleSholat(s.DB, c)

	fmt.Println(deleteJs)

	responseMessage := "Succes"
	if err != nil {
		responseMessage = "tidak ada ID"
	}

	response.GenericJsonResponse(c, http.StatusOK,
		responseMessage,
		nil)
}
