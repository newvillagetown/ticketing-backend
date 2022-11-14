package common

import (
	"fmt"
	"main/common/aws"
	"main/common/env"
	"main/common/oauth/google"
)

func InitServer() error {
	if err := env.InitEnv(); err != nil {
		fmt.Sprintf("서버 에러 발생 : %s", err.Error())
		return err
	}
	if err := aws.InitAws(); err != nil {
		fmt.Sprintf("aws 초기화 에러 : %s", err.Error())
		return err
	}
	if err := google.GoogleOauthInit(); err != nil {
		fmt.Sprintf("구글 초기화 에러 : %s", err.Error())
		return err
	}
	/*
		if err := mongodb.InitMongoDB(); err != nil {
			fmt.Sprintf("mongoDB 초기화 에러 : %s", err.Error())
			return err
		}
		if err := mysql.InitMySQL(); err != nil {
			fmt.Sprintf("mysql 초기화 에러 : %s", err.Error())
			return err
		}

	*/
	return nil
}
