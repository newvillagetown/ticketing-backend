package envCommon

import (
	"fmt"
	"os"
)

type envStruct struct {
	Port    string
	Project string
	Env     string
	Region  string
	IsLocal bool
}

// Env : Environment
var Env envStruct

// 사용하는 환경 변수 네임 설정 함수
func InitVarNames() []string {
	result := make([]string, 0)
	result = append(result, "PORT")
	result = append(result, "PROJECT")
	result = append(result, "ENV")
	result = append(result, "REGION")
	result = append(result, "IS_LOCAL")
	return result
}

// 사용할 환경 변수 값들 초기화해주는 함수
func InitEnv() error {
	envVarNames := InitVarNames()
	envs, err := getOSLookupEnv(envVarNames)
	if err != nil {
		return err
	}
	Env = envStruct{
		Port:    envs["PORT"],
		Project: envs["PROJECT"],
		Env:     envs["ENV"],
		Region:  envs["REGION"],
		IsLocal: envIsLocal(envs["IS_LOCAL"]),
	}
	return nil
}
func envIsLocal(isLocal string) bool {
	if isLocal != "true" {
		return false
	} else {
		return true
	}
}

func PointTrue() *bool {
	result := true
	return &result
}
func PointFalse() *bool {
	result := false
	return &result
}
func getOSLookupEnv(envVarNames []string) (map[string]string, error) {
	result := map[string]string{}
	var ok bool
	for _, envVarName := range envVarNames {
		if result[envVarName], ok = os.LookupEnv(envVarName); !ok {
			return nil, fmt.Errorf("os lookup get failed")
		}
	}
	return result, nil
}
