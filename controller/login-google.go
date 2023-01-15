package controller

import (
	"encoding/json"
	"fmt"
	"ilmi_backend/auth"
	formaterror "ilmi_backend/format-error"
	"ilmi_backend/models"
	"ilmi_backend/response"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
)

func (s *Server) HandleHomeSigIn(w http.ResponseWriter, r *http.Request) {
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
        <a href="/api/v1/oauth2/SignInOauth2" class="btn btn-danger"><span class="fa fa-google"></span> SignUp with Google</a>
    </div>
</div>
</body>
</html>`
	fmt.Fprint(w, html)
}

func (s *Server) SignInOauth2(w http.ResponseWriter, r *http.Request) {
	urlGmail := s.InitializeOauthGoogleConfig().AuthCodeURL(randomState)
	http.Redirect(w, r, urlGmail, http.StatusTemporaryRedirect)
}

func (s *Server) CallBackSignInOauth2(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	if r.FormValue("state") != randomState {
		fmt.Print("state is not valid")
		// http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		response.ERROR(w, http.StatusTemporaryRedirect, nil, "")
		return
	}

	token, err := s.InitializeOauthGoogleConfig().Exchange(oauth2.NoContext, r.FormValue("code"))
	if err != nil {
		fmt.Printf("could not get token : %s\n", err.Error())
		response.ERROR(w, http.StatusTemporaryRedirect, err, "")
		return

	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v1/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Printf("could not create get reques : %s\n", err.Error())
		response.ERROR(w, http.StatusTemporaryRedirect, err, "")
		return
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	var jsonData ResponseGmail
	json.Unmarshal(respBody, &jsonData)
	if err != nil {
		fmt.Printf("could not parse response : %s\n", err.Error())
		response.ERROR(w, http.StatusTemporaryRedirect, err, "")
		return
	}

	loginOauth, err := s.ChekEmailOauth2(jsonData.Email)

	updateTokenUser, err := s.UpdateTokenUser(jsonData.Email, loginOauth)

	fmt.Print(updateTokenUser)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		response.ERROR(w, http.StatusInternalServerError, formattedError, "")
		return
	}

	user.PrepareUser()

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, loginOauth))
	fmt.Fprintf(w, "Response: %s", respBody)
}

// chek email
func (s *Server) ChekEmailOauth2(email string) (string, error) {
	user := models.User{}
	err := s.DB.Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	return auth.CreateToken(user.Id)
}