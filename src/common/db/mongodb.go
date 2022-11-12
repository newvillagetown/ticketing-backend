package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"main/common"
	"time"
)

func InitMongoDB() error {

	connInfos, err := GetEnvMongoDB()
	connURI := MakeMongoDBConnURI(connInfos)
	client, err := mongo.NewClient(options.Client().ApplyURI(connURI))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	defer client.Disconnect(ctx)
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return err
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(databases)
	return err
}

func GetEnvMongoDB() ([]string, error) {

	connInfos, err := common.AwsGetParams([]string{
		fmt.Sprintf("%s-%s-mongodb-id", common.Env.Env, common.Env.Project),
		fmt.Sprintf("%s-%s-mongodb-pw", common.Env.Env, common.Env.Project),
		fmt.Sprintf("%s-%s-mongodb-uri", common.Env.Env, common.Env.Project),
	})
	if err != nil {
		return nil, err
	}
	return connInfos, nil
}

func MakeMongoDBConnURI(connInfos []string) string {
	database := common.Env.Env + "_" + common.Env.Project
	result := fmt.Sprintf("mongodb://%s:%s@%s/%s?retryWrites=true&w=majority&authSource=admin", connInfos[0], connInfos[1], connInfos[2], database)

	return result
}
