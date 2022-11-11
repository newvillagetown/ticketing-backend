package common

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	AwsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func InitAws() error {
	var awsConfig aws.Config
	var err error

	if Env.IsLocal {
		awsConfig, err = AwsConfig.LoadDefaultConfig(context.TODO(),
			AwsConfig.WithRegion(Env.Region),
			AwsConfig.WithSharedConfigProfile("breathings"))
		if err != nil {
			return fmt.Errorf("init aws - region : %s / profile : breathings", Env.Region)
		}
	} else {
		awsConfig, err = AwsConfig.LoadDefaultConfig(context.TODO(),
			AwsConfig.WithRegion(Env.Region))
		if err != nil {
			return fmt.Errorf("init aws - region : %s", Env.Region)
		}
	}

	AwsClientSsm = ssm.NewFromConfig(awsConfig)

	return nil
}
