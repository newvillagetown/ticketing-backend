package google

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
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
	/*
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
	*/
	clientID := "850218569266-agoojlqbo5ffuvmb0t6m0j431pl3p5v8.apps.googleusercontent.com"
	callbackURL := "http://demo-alb-574657214.us-east-1.elb.amazonaws.com/v0.1/auth/google/signin/callback"
	clientSecret := "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"
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
