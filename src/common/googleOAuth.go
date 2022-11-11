package common

import (
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
type webAuthMeta struct {
	GoogleClientID  string `json:"googleClientID" validate:"required"`
	GoogleSecretKey string `json:"googleSecretKey" validate:"required"`
	Hd              string `json:"hd",validate:"required"`
}

var OAuthConf *oauth2.Config

// TODO 추후 config 정보는 aws ssm에서 관리할 예정
const (
	CallBackURL         = "http://localhost:3000/v0.1/auth/google/signin/callback"
	UserInfoAPIEndpoint = "https://www.googleapis.com/oauth2/v3/userinfo"
	ScopeEmail          = "https://www.googleapis.com/auth/userinfo.email"
	ScopeProfile        = "https://www.googleapis.com/auth/userinfo.profile"
	ClientID            = "850218569266-agoojlqbo5ffuvmb0t6m0j431pl3p5v8.apps.googleusercontent.com"
	ClientSecret        = "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"
)

func GoogleOauthInit() error {
	OAuthConf = &oauth2.Config{
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		RedirectURL:  CallBackURL,
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
