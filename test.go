package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"log"
	"time"
)

type Person struct {
	Name string
	Age  int
	City string
}

func main()  {
	//var client *mongo.Client
	//var mongoCtx context.Context
	//
	//for{
	//	logrus.Printf("Connecting to MongoDB...")
	//	mongoCtx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	//	db, err := mongo.Connect(mongoCtx, options.Client().ApplyURI(fmt.Sprintf("mongodb://root:123456@localhost:%v", 27017)))
	//	if err!=nil{
	//		logrus.Printf("Could not connect to MongoDB: %v\n", err.Error())
	//		break
	//	}else {
	//		logrus.Printf("Connected to Mongodb successfully")
	//		client = db
	//		break
	//	}
	//}
	//
	//collection := client.Database("mydb").Collection("persons")
	//ruan := Person{"Ruan", 34, "Cape Town"}
	//insertResult, err := collection.InsertOne(context.TODO(), ruan)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)


	var client *mongo.Client
	var mongoCtx context.Context
	var err error

	for{
		// Set client options
		clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://root:123456@localhost:%v", 27017))
		mongoCtx, _ = context.WithTimeout(context.Background(), 2*time.Second)

		// Connect to MongoDB
		client, err = mongo.Connect(mongoCtx, clientOptions)
		if err != nil {
			log.Fatal("111 : ", err.Error())
			break
		}

		// Check the connection
		err = client.Ping(mongoCtx, nil)
		if err != nil {
			log.Fatal("222 : ", err.Error())
			break
		}

		fmt.Println("Connected to MongoDB!")
		break
	}
}
