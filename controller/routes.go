package controller

func (s *Server) InitializeRoutes() {
	// Api Is Active
	//s.Router.GET("/", (s.DefaultServer)).Methods("GET")

	v1 := s.Router.Group("/api/v1")
	{
		v1.GET("/status", s.DefaultServer)

		v1.GET("/oauth2/homeRegister", s.HandleHomeSigUp)
		v1.GET("/oauth2/registerAuth2", s.SignUpOauth2)
		v1.GET("/oauth2/callbackRegisterAuth2", s.CallBackSignUpOauth2)

		v1.GET("/oauth2/homeLogin", s.HandleHomeSigIn)
		v1.GET("/oauth2/loginAuth2", s.SignInOauth2)
		v1.GET("/oauth2/callbackAuth2", s.CallBackSignInOauth2)

		v1.POST("/register", s.CreateUser)
		v1.POST("/login", s.Login)

		v1.POST("/sholat/write-sholat-history", s.CreateHistorySholat)
		v1.GET("/sholat/get-sholat-history", s.GetHistorySholatByDate)

		v1.GET("/user/:id", s.GetProfile)

		//forgor password
		v1.POST("/chekEmail", s.ChekEmailOtp)
		v1.POST("/validateOtp", s.ValidateOtp)
		v1.PUT("/changePassword/:id", s.ChangePassword)
	}
}
