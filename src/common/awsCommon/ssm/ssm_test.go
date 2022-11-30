package ssm

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	AwsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAwsGetParam(t *testing.T) {
	t.Run("aws ssm get param", func(t *testing.T) {
		var awsConfig aws.Config
		var err error

		awsConfig, _ = AwsConfig.LoadDefaultConfig(context.TODO(),
			AwsConfig.WithRegion("us-east-1"),
			AwsConfig.WithSharedConfigProfile("breathings"))
		AwsClientSsm = ssm.NewFromConfig(awsConfig)
		got, err := AwsGetParam("dev-ticketing-test")
		assert.Nil(t, err)
		want := "aAScx2123aCs1DasdPs"
		assert.Equal(t, got, want)
	})
}

func TestAwsGetParams(t *testing.T) {
	t.Run("aws ssm get params", func(t *testing.T) {
		var awsConfig aws.Config
		var err error

		awsConfig, _ = AwsConfig.LoadDefaultConfig(context.TODO(),
			AwsConfig.WithRegion("us-east-1"),
			AwsConfig.WithSharedConfigProfile("breathings"))
		AwsClientSsm = ssm.NewFromConfig(awsConfig)
		got, err := AwsGetParams([]string{"dev-ticketing-test", "dev-ticketing-test2"})
		assert.Nil(t, err)
		want := []string{"aAScx2123aCs1DasdPs", "cadmSD1DCdc12995C"}
		assert.Equal(t, got, want)
	})
}
