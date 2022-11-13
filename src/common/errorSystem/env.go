package errorSystem

import "net/http"

// 에러 타입을 구분
type ErrType string

// 에러가 어디서 발생했는지 확인용
type IErrFrom string

const (
	ErrFromClient   = IErrFrom("client")
	ErrFromInternal = IErrFrom("internal")
	ErrFromMongoDB  = IErrFrom("mongoDB")
	ErrFromAws      = IErrFrom("aws")
	ErrFromAwsS3    = IErrFrom("aws_s3")
	ErrFromAwsSsm   = IErrFrom("aws_ssm")
)

const (
	ErrBadParameter         = ErrType("PARAM_BAD")
	ErrAuthFailed           = ErrType("AUTH_FAILED")
	ErrAuthInActive         = ErrType("AUTH_INACTIVE")
	ErrUserNotExisted       = ErrType("USR_NOT_EXISTED")
	ErrUserAlreadyExisted   = ErrType("USR_ALREADY_EXISTED")
	ErrBadToken             = ErrType("TOKEN_BAD")
	ErrAuthPolicyViolation  = ErrType("POLICY_VIOLATION")
	ErrInternalServer       = ErrType("INTERNAL_SERVER")
	ErrInternalDB           = ErrType("INTERNAL_DB")
	ErrPartner              = ErrType("PARTNER")
	ErrNotMatchedLoginInfo  = ErrType("NOT_MATCHED_LOGIN_INFO")
	ErrNotMatchedSignupInfo = ErrType("NOT_MATCHED_SIGNUP_INFO")
	ErrInvalidAuthCode      = ErrType("INVALID_AUTH_CODE")
	ErrExpiredAuthCode      = ErrType("EXPIRED_AUTH_CODE")
)

// 에러 타입에 따라서 httpCode 맵핑
var ErrHttpCode = map[string]int{
	"PARAM_BAD":               http.StatusBadRequest,
	"AUTH_FAILED":             http.StatusUnauthorized,
	"AUTH_INACTIVE":           http.StatusForbidden,
	"USR_NOT_EXISTED":         http.StatusBadRequest,
	"USR_ALREADY_EXISTED":     http.StatusBadRequest,
	"TOKEN_BAD":               http.StatusUnauthorized,
	"POLICY_VIOLATION":        http.StatusUnauthorized,
	"INTERNAL_SERVER":         http.StatusInternalServerError,
	"INTERNAL_DB":             http.StatusInternalServerError,
	"PARTNER":                 http.StatusInternalServerError,
	"NOT_MATCHED_LOGIN_INFO":  http.StatusBadRequest,
	"NOT_MATCHED_SIGNUP_INFO": http.StatusBadRequest,
	"INVALID_AUTH_CODE":       http.StatusBadRequest,
	"EXPIRED_AUTH_CODE":       http.StatusBadRequest,
}
