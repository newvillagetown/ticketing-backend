package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"main/common/aws/ssm"
	"main/common/env"
)

func InitMySQL() error {
	connInfos, err := GetEnvMySQL()
	connURI := MakeMySQLConnURI(connInfos)
	db, err := sql.Open("mysql", connURI)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	return nil
}

func GetEnvMySQL() ([]string, error) {

	connInfos, err := ssm.AwsGetParams([]string{
		fmt.Sprintf("%s-%s-mysql-id", env.Env.Env, env.Env.Project),
		fmt.Sprintf("%s-%s-mysql-pw", env.Env.Env, env.Env.Project),
		fmt.Sprintf("%s-%s-mysql-uri", env.Env.Env, env.Env.Project),
	})
	if err != nil {
		return nil, err
	}
	return connInfos, nil
}

func MakeMySQLConnURI(connInfos []string) string {
	database := env.Env.Env + "_" + env.Env.Project
	result := fmt.Sprintf("%s:%s@tcp(%s)/%s", connInfos[0], connInfos[1], connInfos[2], database)

	return result
}
