package mysqlCommon

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main/common/awsCommon/ssm"
	"main/common/envCommon"
	"time"
)

var MysqlDB *sql.DB
var GormDB *gorm.DB

func InitMySQL() error {
	connInfos, err := GetEnvMySQL()
	connURI := MakeMySQLConnURI(connInfos)
	MysqlDB, err = sql.Open("mysql", connURI)
	if err != nil {
		return err
	}
	err = MysqlDB.Ping()
	if err != nil {
		return err
	}
	/*
		GORM perform write (create/update/delete) operations run inside a transaction to ensure data consistency,
		you can disable it during initialization if it is not required, you will gain about 30%+ performance improvement after that
	*/
	GormDB, err = gorm.Open(mysql.New(mysql.Config{
		Conn: MysqlDB,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return err
	}

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
	if envCommon.Env.IsLocal == true {
		url, err := ssm.AwsGetParam(fmt.Sprintf("%s-%s-mysql-uri-local", envCommon.Env.Env, envCommon.Env.Project))
		if err != nil {
			return nil, err
		}
		connInfos[2] = url
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

func NowDateGenerate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func EpochToTime(t int64) time.Time {
	return time.Unix(t, t%1000*1000000)
}
func EpochToTimeString(t int64) string {
	return time.Unix(t, t%1000*1000000).String()
}

func TimeStringToEpoch(t string) int64 {
	date, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", t)
	return date.Unix()
}
