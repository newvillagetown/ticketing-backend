package google

import (
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"main/common/aws/ssm"
	"main/common/env"
)

var OAuthConf *oauth2.Config

// TODO 추후 config 정보는 aws ssm에서 관리할 예정
const (
	UserInfoAPIEndpoint = "https://www.googleapis.com/oauth2/v3/userinfo"
	ScopeEmail          = "https://www.googleapis.com/auth/userinfo.email"
	ScopeProfile        = "https://www.googleapis.com/auth/userinfo.profile"
)

func GoogleOauthInit() error {
	clientID, err := ssm.AwsGetParam(fmt.Sprintf("%s-%s-google-oauth-clientid", env.Env.Env, env.Env.Project))
	if err != nil {
		return err
	}
	clientSecret, err := ssm.AwsGetParam(fmt.Sprintf("%s-%s-google-oauth-clientsecret", env.Env.Env, env.Env.Project))
	if err != nil {
		return err
	}
	callbackURL, err := ssm.AwsGetParam(fmt.Sprintf("%s-%s-google-oauth-callbackurl", env.Env.Env, env.Env.Project))
	if err != nil {
		return err
	}
	if env.Env.IsLocal == true {
		callbackURL = "http://localhost:3000/v0.1/auth/google/signin/callback"
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
