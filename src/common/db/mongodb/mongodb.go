package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"main/common/aws/ssm"
	"main/common/env"
	"time"
)

var MongoClient *mongo.Client

func InitMongoDB() error {
	connInfos, err := GetEnvMongoDB()
	if err != nil {
		return err
	}
	connURI := MakeMongoDBConnURI(connInfos)

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
		fmt.Sprintf("%s-%s-mongodb-id", env.Env.Env, env.Env.Project),
		fmt.Sprintf("%s-%s-mongodb-pw", env.Env.Env, env.Env.Project),
		fmt.Sprintf("%s-%s-mongodb-uri", env.Env.Env, env.Env.Project),
	})
	if err != nil {
		return nil, err
	}
	return connInfos, nil
}

func MakeMongoDBConnURI(connInfos []string) string {
	database := env.Env.Env + "_" + env.Env.Project
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
