package controller

func (s *Server) InitializeRoutes() {
	// Api Is Active
	//s.Router.GET("/", (s.DefaultServer)).Methods("GET")

	v1 := s.Router.Group("/api/v1")
	{
		v1.POST("/register", s.CreateUser)
		v1.POST("/login", s.Login)

		v1.POST("/sholat/write-sholat-history", s.CreateHistorySholat)
		v1.GET("/sholat/get-sholat-history", s.GetHistorySholatByDate)

	}
}
