package controller

import "os"

func (s *Server) InitializeRoutes() {
	// Api Is Active
	s.Router.HandleFunc("/", (s.DefaultServer)).Methods("GET")

	s.Router.HandleFunc(os.Getenv("BASE_URL")+"register", s.CreateUser).Methods("POST")
	s.Router.HandleFunc(os.Getenv("BASE_URL")+"login", s.Login).Methods("POST")

	//oauth sigup
	s.Router.HandleFunc(os.Getenv("BASE_URL")+"oauth2/home", s.HandleHomeSigUp)
	s.Router.HandleFunc(os.Getenv("BASE_URL")+"oauth2/SignUpOauth2", s.SignUpOauth2).Methods("POST")
	s.Router.HandleFunc(os.Getenv("BASE_URL")+"oauth2/callback", s.CallBackSignUpOauth2)

}
