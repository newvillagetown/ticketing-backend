package ssm

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"main/common/envCommon"
	"strings"
)

var AwsClientSsm *ssm.Client

func AwsGetParam(path string) (string, error) {
	ctx := context.TODO()
	param, err := AwsClientSsm.GetParameter(ctx, &ssm.GetParameterInput{
		Name:           aws.String(path),
		WithDecryption: envCommon.PointTrue(),
	})
	if err != nil {
		return "", fmt.Errorf("get ssm param : %s", path)
	}
	return aws.ToString(param.Parameter.Value), nil
}

func AwsGetParams(paths []string) ([]string, error) {
	ctx := context.TODO()
	params, err := AwsClientSsm.GetParameters(ctx, &ssm.GetParametersInput{
		Names:          paths,
		WithDecryption: envCommon.PointTrue(),
	})
	if err != nil {
		return nil, fmt.Errorf("get ssm params : %s", strings.Join(paths, ","))
	}
	result := make([]string, len(paths))
	for i, path := range paths {
		val := ""
		for _, parameter := range params.Parameters {
			if strings.Contains(aws.ToString(parameter.ARN), path) {
				val = aws.ToString(parameter.Value)
				break
			}
		}
		result[i] = val
	}
	return result, nil
}
