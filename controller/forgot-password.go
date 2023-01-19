package controller

import (
	"fmt"
	"ilmi_backend/helper"
	"ilmi_backend/models"
	"ilmi_backend/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

// chek email & auto update OTP
func (s *Server) ChekEmailOtp(c *gin.Context) {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	user := models.User{}

	if err := c.BindJSON(&user); err != nil {
		response.GenericJsonResponse(c, http.StatusBadRequest,
			err.Error(),
			nil)
		return
	}

	checkEmail, err := user.ChekEmail(s.DB, user.Email)
	fmt.Println(checkEmail)

	if err != nil {
		response.GenericJsonResponse(c, http.StatusBadGateway, "Email Tidak Terdaftar", "null")
	} else {
		newOtp := helper.AutoGenerateOtp(int(user.Id))
		println("idUser : ", user.Id)

		err = user.UpdateOtp(s.DB, user.Email, newOtp)
		if err != nil {
			response.ErrorResponse(c, http.StatusBadRequest, err)
			return
		}

		emailUser := user.Email
		toEmailUserSendEmail := []string{emailUser}

		otpToStr := strconv.FormatUint(uint64(newOtp), 10)
		//otp send gmail
		subject := "IlmiApp"
		otpMessageToEmail := "OTP : " + otpToStr
		senderName := "PT.Testing"

		helper.SendBodyOtpToEmail(toEmailUserSendEmail, senderName, subject, otpMessageToEmail)
		response.GenericJsonResponse(c, http.StatusOK, "Succes", otpToStr)
	}
}

// validation Otp
func (s *Server) ValidateOtp(c *gin.Context) {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	user := models.User{}
	if err := c.BindJSON(&user); err != nil {
		response.GenericJsonResponse(c, http.StatusBadRequest,
			err.Error(),
			nil)
		return
	} else {
		chekOtp, err := user.ChekOtp(s.DB, c, uint64(user.Otp))
		if err != nil {
			response.GenericJsonResponse(c, http.StatusBadRequest, "Not Otp", "null")
		}
		response.GenericJsonResponse(c, http.StatusOK, "Validate Succes", chekOtp.Otp)
	}
}

// change Password
func (s *Server) ChangePassword(c *gin.Context) {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	user := models.User{}

	if err := c.BindJSON(&user); err != nil {
		response.GenericJsonResponse(c, http.StatusBadRequest,
			err.Error(),
			nil)
		return
	}

	hashedPassword := models.HashPasswordToSha256(user.Password)

	changePassword, err := user.ResertPassword(s.DB, c, hashedPassword)
	print(changePassword)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
	} else {
		response.GenericJsonResponse(c, http.StatusOK, "Succes", hashedPassword)
	}

}
