package firebaseCommon

import (
	"context"
	"firebase.google.com/go"
	"fmt"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
	"main/common/awsCommon/ssm"
	"main/common/envCommon"
	"time"
)

type UserData struct {
	Email      string
	UID        string
	LastSignin time.Time
	Created    time.Time
}

type FirebaseCredential struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyId            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthUri                 string `json:"AuthUri"`
	TokenUri                string `json:"token_uri"`
	AuthProviderX509CertUrl string `json:"auth_provider_x509_cert_url"`
	ClientX509CertUrl       string `json:"client_x509_cert_url"`
}

func Init() error {
	// s3 에 보관하고 실행할때마다 파일을 가져온다.
	users, err := getAllUsers()
	if err != nil {
		return err
	}
	fmt.Println(users)
	return nil
}

// 유저 가져와서 업데이트 치기
// 가져올 수 있는 데이터가 이메일, 생성 날짜, 마지막 로그인 날짜, 사용자UID
// 매일 자정에 배치 한번 돌면 됨.
func getAllUsers() ([]UserData, error) {
	// Initialize the Firebase Admin SDK
	ctx := context.TODO()
	firebaseInfo, err := ssm.AwsGetParam("dev_firebase_credential")
	sa := option.WithCredentialsJSON([]byte(firebaseInfo))

	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase app: %v", err)
	}

	// Get a client for the Firebase Auth service
	authService, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed to get Firebase Auth client: %v", err)
	}
	// 어제날짜 추출
	now := time.Now()
	yesterday := now.Add(-24 * time.Hour)
	KST, _ := time.LoadLocation("Local")
	startTime := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, KST).Local() // 자정
	endTime := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 24, 0, 0, 0, KST).Local()  // 자정 직전

	users := make([]UserData, 0)
	iter := authService.Users(ctx, "")
	for {
		user, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		if envCommon.TimeToEpochMillis(startTime) <= user.UserMetadata.LastLogInTimestamp && user.UserMetadata.LastLogInTimestamp < envCommon.TimeToEpochMillis(endTime) {
			lastSignin := envCommon.EpochToTimeMillis(user.UserMetadata.LastLogInTimestamp)
			created := envCommon.EpochToTimeMillis(user.UserMetadata.CreationTimestamp)
			users = append(users, UserData{UID: user.UID, Email: user.Email, LastSignin: lastSignin, Created: created})
		}
	}
	return users, nil
}
