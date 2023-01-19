package controller

import (
	"ilmi_backend/models"
	"ilmi_backend/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func (s *Server) GetAllProvinsi(c *gin.Context) {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	provinsi := models.Provinsi{}
	page := c.Query("page")
	offsets := c.Query("offset")
	
	err := provinsi.PagingValidate(page, offsets)

	getAllKota, _ := provinsi.GetAllProvinsi(s.DB, page, offsets)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest,
			err)
	} else {
		responseMessage := "Succes"

		response.GenericJsonResponse(c, http.StatusOK,
			responseMessage,
			getAllKota)
	}

}
