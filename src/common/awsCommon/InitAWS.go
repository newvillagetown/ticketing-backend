package awsCommon

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	AwsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	ssm2 "main/common/awsCommon/ssm"
	"main/common/envCommon"
)

func InitAws() error {
	var awsConfig aws.Config
	var err error

	awsConfig, err = AwsConfig.LoadDefaultConfig(context.TODO(),
		AwsConfig.WithRegion(envCommon.Env.Region),
		AwsConfig.WithSharedConfigProfile("breathings"))
	if err != nil {
		return fmt.Errorf("init aws - region : %s / profile : breathings", envCommon.Env.Region)
	}

	ssm2.AwsClientSsm = ssm.NewFromConfig(awsConfig)

	return nil
}
