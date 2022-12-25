package noticeCommon

import (
	"fmt"
	"main/common/awsCommon/ssm"
	"main/common/envCommon"
)

var GoogleChatUrl string

func GoogleChatInit() error {
	var err error
	GoogleChatUrl, err = ssm.AwsGetParam(fmt.Sprintf("%s-%s-google-webhook-url", envCommon.Env.Env, envCommon.Env.Project))
	if err != nil {
		return err
	}
	return nil
}
