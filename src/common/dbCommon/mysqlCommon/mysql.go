package mysqlCommon

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"main/common/awsCommon/ssm"
	"main/common/envCommon"
)

var MysqlDB *sql.DB

func InitMySQL() error {
	connInfos, err := GetEnvMySQL()
	connURI := MakeMySQLConnURI(connInfos)
	MysqlDB, err = sql.Open("mysql", connURI)
	if err != nil {
		fmt.Println(err)
		return err
	}
	//	defer MysqlDB.Close()

	return nil
}

func GetEnvMySQL() ([]string, error) {

	connInfos, err := ssm.AwsGetParams([]string{
		fmt.Sprintf("%s-%s-mysql-id", envCommon.Env.Env, envCommon.Env.Project),
		fmt.Sprintf("%s-%s-mysql-pw", envCommon.Env.Env, envCommon.Env.Project),
		fmt.Sprintf("%s-%s-mysql-uri", envCommon.Env.Env, envCommon.Env.Project),
	})
	if err != nil {
		return nil, err
	}
	return connInfos, nil
}

func MakeMySQLConnURI(connInfos []string) string {
	database := envCommon.Env.Env + "_" + envCommon.Env.Project
	result := fmt.Sprintf("%s:%s@tcp(%s)/%s", connInfos[0], connInfos[1], connInfos[2], database)

	return result
}

func PKIDGenerate() string {
	//uuid 로 생성
	result := (uuid.New()).String()
	return result
}
