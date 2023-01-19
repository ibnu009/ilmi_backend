package controller

import (
	"encoding/json"
	"fmt"
	formaterror "ilmi_backend/format-error"
	"ilmi_backend/models"
	"ilmi_backend/response"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func (s *Server) HandleHomeSigIn(c *gin.Context) {
	var html = `
	<head>
    <title>Google SignIn</title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css">
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
    <style>
        body        { padding-top:70px; }
    </style>
</head>
<body>
<div class="container">
    <div class="jumbotron text-center text-success">
        <h1><span class="fa fa-lock"></span> Social Authentication</h1>
        <p>Login or Login Google with:</p>
        <a href="/api/v1/oauth2/loginAuth2" class="btn btn-danger"><span class="fa fa-google"></span> SignUp with Google</a>
    </div>
</div>
</body>
</html>`
	fmt.Fprint(c.Writer, html)
}

func (s *Server) SignInOauth2(c *gin.Context) {
	urlGmail := s.InitializeOauthGoogleConfig().AuthCodeURL(randomState)
	http.Redirect(c.Writer, c.Request, urlGmail, http.StatusTemporaryRedirect)
}

func (s *Server) CallBackSignInOauth2(c *gin.Context) {
	user := models.User{}
	if c.Request.FormValue("state") != randomState {
		fmt.Print("state is not valid")
		response.ErrorResponse(c, http.StatusTemporaryRedirect, nil)
		return
	}

	token, err := s.InitializeOauthGoogleConfig().Exchange(oauth2.NoContext, c.Request.FormValue("code"))
	if err != nil {
		fmt.Printf("could not get token : %s\n", err.Error())
		return

	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v1/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Printf("could not create get reques : %s\n", err.Error())
		response.ErrorResponse(c, http.StatusTemporaryRedirect, err)
		return
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	var jsonData ResponseGmail

	json.Unmarshal(respBody, &jsonData)
	if err != nil {
		fmt.Printf("could not parse response : %s\n", err.Error())
		response.ErrorResponse(c, http.StatusTemporaryRedirect, err)
		return
	}

	loginOauth, err := user.ChekEmailOauth2(s.DB, jsonData.Email)
	if err != nil {
		response.GenericJsonResponse(c, http.StatusBadRequest,
			"Maaf Email Anda Belum Terdaftar",
			nil)
		return
	}

	updateTokenUser, err := s.UpdateTokenUser(jsonData.Email, loginOauth)

	fmt.Print(updateTokenUser)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		response.ErrorResponse(c, http.StatusInternalServerError, formattedError)
		return
	}

	user.PrepareUser()

	c.Writer.Header().Set("Location", fmt.Sprintf("%s%s/%d", c.Request.Host, c.Request.RequestURI, loginOauth))
	response.GenericJsonResponse(
		c, http.StatusOK, "Succes Login", updateTokenUser,
	)
}
