package controller

import (
	"fmt"
	"ilmi_backend/models"
	"ilmi_backend/response"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

// config to google apis
func GoogleOauthConfig(redirectUrl, clienId, clientSecret, scopes string) (*oauth2.Config, error) {
	var err error
	googleOauthConfig := &oauth2.Config{
		RedirectURL:  redirectUrl,
		ClientID:     clienId,
		ClientSecret: clientSecret,
		Scopes:       []string{scopes},
		Endpoint:     google.Endpoint,
	}
	if googleOauthConfig != nil {
		return googleOauthConfig, nil
	}
	return nil, err
}

func (s *Server) InitializeOauthGoogleConfig() *oauth2.Config {
	configGoolgeApis, err := GoogleOauthConfig(os.Getenv("REDIRECT_URL"), os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_SECRET"), os.Getenv("GOOGLE_SCOPES"))
	if err != nil {
		fmt.Println(err)
	}
	return configGoolgeApis
}

// config to db local
func (s *Server) InitializeServer(DbDriver, DbHost, DbUser, DbPassword, DbName, DbPort string) {
	var err error
	if DbDriver == "mysql" {
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		s.DB, err = gorm.Open(DbDriver, dsn)
		if err != nil {
			fmt.Print("not connected to database", DbDriver)
			log.Fatal("This The Error :", err)
		} else {
			fmt.Printf("connected to the %v database", DbDriver)
		}
	}
	gorm.DefaultCallback.Create().Remove("mssql:set_identity_insert")
	// migrate
	s.DB.AutoMigrate(
		models.User{},
		models.HistorySholat{},
	)
	s.Router = gin.Default()
	s.InitializeRoutes()
}

func (s *Server) DefaultServer(c *gin.Context) {
	response.GenericJsonResponse(c, http.StatusOK, "Backend Service Ilmi Is Running", nil)
}

func (s *Server) RunServer(addr string) {
	fmt.Println("Listen Of Port Server : " + os.Getenv("PORT"))
	handler := s.Router
	log.Fatal(http.ListenAndServe(addr, handler))
}
