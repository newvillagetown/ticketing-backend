package awsCommon

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	AwsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/cloudfront/sign"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	ssmCommon "main/common/awsCommon/ssm"
	"main/common/envCommon"
	"main/common/s3Common"
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

	ssmCommon.AwsClientSsm = ssm.NewFromConfig(awsConfig)
	s3Common.AwsClientS3 = s3.NewFromConfig(awsConfig)
	s3Common.AwsClientS3Uploader = manager.NewUploader(s3Common.AwsClientS3)
	s3Common.AwsClientS3Downloader = manager.NewDownloader(s3Common.AwsClientS3)
	s3SignKey, err := ssmCommon.AwsGetParams([]string{
		"cloudfront_signkey_id",
		"cloudfront_signkey_private"})
	if err != nil {
		return err
	}
	signPublicKey := s3SignKey[0]
	block, _ := pem.Decode([]byte(s3SignKey[1]))
	signPrivateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return fmt.Errorf("init aws - sign private key : %s", s3SignKey[1])
	}
	s3Common.AwsS3Signer = sign.NewURLSigner(signPublicKey, signPrivateKey)
	return nil
}
