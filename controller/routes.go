package controller

import (
	"ilmi_backend/middleware"
)

func (s *Server) InitializeRoutes() {
	v1 := s.Router.Group("/api/v1")
	{

		v1.GET("/status", s.StatusServer)

		/*Login Oauth2*/
		v1.GET("/oauth2/homeRegister", s.HandleHomeSigUp)
		v1.GET("/oauth2/registerAuth2", s.SignUpOauth2)
		v1.GET("/oauth2/callbackRegisterAuth2", s.CallBackSignUpOauth2)

		/*Register Oauth2*/
		v1.GET("/oauth2/homeLogin", s.HandleHomeSigIn)
		v1.GET("/oauth2/loginAuth2", s.SignInOauth2)
		v1.GET("/oauth2/callbackAuth2", s.CallBackSignInOauth2)

		/*Login & Register*/
		v1.POST("/register", s.CreateUser)
		v1.POST("/login", s.Login)

		/*User*/
		v1.GET("/user", s.GetProfile)

		/*forgor password*/
		v1.POST("/chekEmail", s.ChekEmailOtp)
		v1.POST("/validateOtp", s.ValidateOtp)
		v1.PUT("/changePassword/:id", s.ChangePassword)

		/*History Sholat*/
		v1.POST("/sholat/write-sholat-history", s.CreateHistorySholat)
		v1.GET("/sholat/get-sholat-history", s.GetHistorySholatByDate)

		/*Jadwal Sholat*/
		v1.POST("/schedule-sholat/created", s.CreatedScheduleSholat)
		v1.PUT("/schedule-sholat/update/:id", s.UpdateScheduleSholat)
		v1.GET("/schedule-sholat", middleware.SetMiddlewareAuthentication(s.GetScheduleSholat))
		v1.DELETE("/schedule-sholat/delete/:id", middleware.SetMiddlewareAuthentication(s.DeleteScheduleSholat))

		/*Kota*/
		v1.GET("/kota", s.GetAllKota)

		/*Provinsi*/
		v1.GET("/provinsi", s.GetAllProvinsi)
	}
}
