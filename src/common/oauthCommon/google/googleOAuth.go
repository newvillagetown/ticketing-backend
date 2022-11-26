package google

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"main/common/awsCommon/ssm"
	"main/common/envCommon"
)

var OAuthConf *oauth2.Config

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

const (
	UserInfoAPIEndpoint = "https://www.googleapis.com/oauth2/v3/userinfo"
	ScopeEmail          = "https://www.googleapis.com/auth/userinfo.email"
	ScopeProfile        = "https://www.googleapis.com/auth/userinfo.profile"
)

func GoogleOauthInit() error {
	clientID, err := ssm.AwsGetParam(fmt.Sprintf("%s-%s-google-oauth-clientid", envCommon.Env.Env, envCommon.Env.Project))
	if err != nil {
		return err
	}
	clientSecret, err := ssm.AwsGetParam(fmt.Sprintf("%s-%s-google-oauth-clientsecret", envCommon.Env.Env, envCommon.Env.Project))
	if err != nil {
		return err
	}
	callbackURL, err := ssm.AwsGetParam(fmt.Sprintf("%s-%s-google-oauth-callbackurl", envCommon.Env.Env, envCommon.Env.Project))
	if err != nil {
		return err
	}

	if envCommon.Env.IsLocal == true {
		callbackURL = fmt.Sprintf("http://localhost:%s/v0.1/auth/google/signin/callback", envCommon.Env.Port)
	}
	OAuthConf = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  callbackURL,
		Scopes:       []string{ScopeEmail, ScopeProfile},
		Endpoint:     google.Endpoint,
	}
	fmt.Println(OAuthConf)
	return nil
}

func GetLoginURL(state string) string {
	return OAuthConf.AuthCodeURL(state)
}

func RandToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
