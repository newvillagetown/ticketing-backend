package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"main/common"
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

	connInfos, err := common.AwsGetParams([]string{
		fmt.Sprintf("%s-%s-mysql-id", common.Env.Env, common.Env.Project),
		fmt.Sprintf("%s-%s-mysql-pw", common.Env.Env, common.Env.Project),
		fmt.Sprintf("%s-%s-mysql-uri", common.Env.Env, common.Env.Project),
	})
	if err != nil {
		return nil, err
	}
	return connInfos, nil
}

func MakeMySQLConnURI(connInfos []string) string {
	database := common.Env.Env + "_" + common.Env.Project
	result := fmt.Sprintf("%s:%s@tcp(%s)/%s", connInfos[0], connInfos[1], connInfos[2], database)

	return result
}
