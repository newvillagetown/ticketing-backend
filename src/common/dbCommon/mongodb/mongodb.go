package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"main/common/awsCommon/ssm"
	"main/common/envCommon"
	"time"
)

var MongoClient *mongo.Client

func InitMongoDB() error {
	connInfos, err := GetEnvMongoDB()
	if err != nil {
		return err
	}
	connURI := MakeMongoDBConnURI(connInfos)
	fmt.Println(connURI)
	err = ConnectMongoDB(connURI)
	if err != nil {
		return err
	}

	//TODO 컬렉션 초기화
	fmt.Println("mongodb connect")

	return err
}

func GetEnvMongoDB() ([]string, error) {
	connInfos, err := ssm.AwsGetParams([]string{
		fmt.Sprintf("%s-%s-mongodb-id", envCommon.Env.Env, envCommon.Env.Project),
		fmt.Sprintf("%s-%s-mongodb-pw", envCommon.Env.Env, envCommon.Env.Project),
		fmt.Sprintf("%s-%s-mongodb-uri", envCommon.Env.Env, envCommon.Env.Project),
	})
	if err != nil {
		return nil, err
	}
	return connInfos, nil
}

func MakeMongoDBConnURI(connInfos []string) string {
	database := envCommon.Env.Env + "_" + envCommon.Env.Project
	result := fmt.Sprintf("mongodb://%s:%s@%s/%s?retryWrites=true&w=majority&authSource=admin", connInfos[0], connInfos[1], connInfos[2], database)
	fmt.Println(result)
	return result
}

func ConnectMongoDB(connURI string) error {
	var err error
	MongoClient, err = mongo.NewClient(options.Client().ApplyURI(connURI))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = MongoClient.Connect(ctx)
	if err != nil {
		return err
	}
	defer MongoClient.Disconnect(ctx)
	err = MongoClient.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return err
	}
	return nil
}
