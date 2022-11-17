package common

import (
	"fmt"
	"main/common/awsCommon"
	"main/common/dbCommon/mysql"
	"main/common/envCommon"
	"main/common/oauthCommon/google"
)

func InitServer() error {
	if err := envCommon.InitEnv(); err != nil {
		fmt.Sprintf("서버 에러 발생 : %s", err.Error())
		return err
	}
	if err := awsCommon.InitAws(); err != nil {
		fmt.Sprintf("aws 초기화 에러 : %s", err.Error())
		return err
	}
	if err := google.GoogleOauthInit(); err != nil {
		fmt.Sprintf("구글 초기화 에러 : %s", err.Error())
		return err
	}
	//TODO 몽고디비 대책이 필요 계속 뻑나네..
	/*
		if err := mongodb.InitMongoDB(); err != nil {
			fmt.Sprintf("mongoDB 초기화 에러 : %s", err.Error())
			return err
		}
	*/
	if err := mysql.InitMySQL(); err != nil {
		fmt.Sprintf("mysql 초기화 에러 : %s", err.Error())
		return err
	}
	return nil
}
