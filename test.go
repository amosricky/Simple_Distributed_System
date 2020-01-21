package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"time"
)

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27041/?connect=direct"))
	if err != nil{
		fmt.Print(1)
		fmt.Print(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil{
		fmt.Print(2)
		fmt.Print(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil{
		fmt.Print(3)
		fmt.Print(err)
	}
	queryCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection := client.Database("myDB").Collection("game")
	cur, curErr := collection.Find(queryCtx, bson.M{})
	if curErr != nil {
		logrus.Warnf(curErr.Error())
	}

	fmt.Print(cur)


	//mongoCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//dbUrl := fmt.Sprintf("mongodb://127.0.0.1:27041")
	//db, err := mongo.Connect(mongoCtx, options.Client().ApplyURI(dbUrl))
	//if err != nil{
	//	fmt.Printf(err.Error())
	//}
	//err = db.Ping(mongoCtx, nil)
	//if err!=nil{
	//	fmt.Printf(err.Error())
	//}else {
	//	fmt.Printf("ok")
	//}
}
