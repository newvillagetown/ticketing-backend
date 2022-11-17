package google

import (
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"main/common/awsCommon/ssm"
	"main/common/envCommon"
)

var OAuthConf *oauth2.Config

// TODO 추후 config 정보는 awsCommon ssm에서 관리할 예정
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
	return nil
}
