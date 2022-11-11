package common

import (
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var OAuthConf *oauth2.Config

// TODO 추후 config 정보는 aws ssm에서 관리할 예정
const (
	UserInfoAPIEndpoint = "https://www.googleapis.com/oauth2/v3/userinfo"
	ScopeEmail          = "https://www.googleapis.com/auth/userinfo.email"
	ScopeProfile        = "https://www.googleapis.com/auth/userinfo.profile"
)

func GoogleOauthInit() error {
	fmt.Println(Env)
	clientID, err := AwsGetParam(fmt.Sprintf("%s-%s-google-oauth-clientid", Env.Env, Env.Project))
	if err != nil {
		return err
	}
	clientSecret, err := AwsGetParam(fmt.Sprintf("%s-%s-google-oauth-clientsecret", Env.Env, Env.Project))
	if err != nil {
		return err
	}
	callbackURL, err := AwsGetParam(fmt.Sprintf("%s-%s-google-oauth-callbackurl", Env.Env, Env.Project))
	if err != nil {
		return err
	}
	fmt.Println(clientID)
	OAuthConf = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  callbackURL,
		Scopes:       []string{ScopeEmail, ScopeProfile},
		Endpoint:     google.Endpoint,
	}
	return nil
}
