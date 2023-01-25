package common

import (
	"fmt"
	"main/common/awsCommon"
	"main/common/dbCommon/mongodbCommon"
	"main/common/dbCommon/mysqlCommon"
	"main/common/envCommon"
	"main/common/nCloudSmsCommon"
	"main/common/noticeCommon"
	"main/common/oauthCommon/google"
	"main/common/pubsubCommon"
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
	if err := pubsubCommon.InitPubSub(); err != nil {
		fmt.Sprintf("pubsub 초기화 에러 : %s", err.Error())
		return err
	}
	if err := nCloudSmsCommon.InitNSms(); err != nil {
		fmt.Sprintf("nCloudSms 초기화 에러 : %s", err.Error())
		return err
	}
	if err := noticeCommon.GoogleChatInit(); err != nil {
		fmt.Sprintf("googleChat 초기화 에러 : %s", err.Error())
		return err
	}

	return nil
}
