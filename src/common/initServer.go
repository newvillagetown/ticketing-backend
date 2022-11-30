package common

import (
	"fmt"
	"main/common/awsCommon"
	"main/common/dbCommon/mongodbCommon"
	"main/common/dbCommon/mysqlCommon"
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
	if err := mongodbCommon.InitMongoDB(); err != nil {
		fmt.Sprintf("mongoDB 초기화 에러 : %s", err.Error())
		return err
	}
	if err := mysqlCommon.InitMySQL(); err != nil {
		fmt.Sprintf("mysqlCommon 초기화 에러 : %s", err.Error())
		return err
	}
	return nil
}
