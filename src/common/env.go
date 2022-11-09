package common

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

var envVarNames = [5]string{"PORT", "PROJECT", "ENV", "REGION", "IS_LOCAL"}

func InitEnv() error {
	var ok bool
	envs := map[string]string{}
	for _, envVarName := range envVarNames {
		if envs[envVarName], ok = os.LookupEnv(envVarName); !ok {
			return fmt.Errorf("init env error")
		}
	}
	fmt.Println(envs)
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
