package main

import (
	"Simple_Distributed_System/pb"
	"Simple_Distributed_System/setting"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

type Person struct {
	Name string
	Age  int
	City string
}

type Person2 struct {
	Name string
	Age  int
}

var db *mongo.Client
var mongoCtx context.Context
type server struct{}

func (s *server) GetScore(ctx context.Context, in *pb.GetScoreRequest) (*pb.GetScoreReply, error){
	logrus.Printf("GetScore request：%s\n", in.Game)
	return &pb.GetScoreReply{Home:[]int32{10,20,30},HomeTotal:60, Visitor:[]int32{30,40,50}, VisitorTotal:120}, nil
}

func (s *server) PutScore(ctx context.Context, in *pb.PutScoreRequest) (*pb.GeneralReply, error) {
	logrus.Printf("PutScore request：%s\n", in)
	return &pb.GeneralReply{Result:"ok"}, nil
}

func (s *server) GetGameList(ctx context.Context, in *pb.GeneralRequest) (*pb.GetGameListReply, error) {
	logrus.Printf("GetGameList request：%s\n", in)
	var result []*pb.GameItem

	for{
		queryCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		collection := db.Database(setting.DatabaseSetting.DBName).Collection(setting.DatabaseSetting.CollectionName)
		cur, err := collection.Find(queryCtx, bson.D{})
		if err != nil {
			logrus.Fatal(err.Error())
			break
		}

		for cur.Next(context.TODO()) {
			Game := &pb.GameItem{}
			err := cur.Decode(Game)
			if err != nil {
				logrus.Fatal(err.Error())
				break
			}
			result = append(result, &Game)
		}
		break
	}
	fmt.Println(result)
	//item1 := pb.GameItem{Id:"123", Game:"123"}
	//item2 := pb.GameItem{Id:"456", Game:"456"}
	GameList := pb.GetGameListReply{}
	return &GameList, nil
	//return &pb.GetGameListReply{result}, nil
}

func (s *server) PostNewGame(ctx context.Context, in *pb.PostNewGameRequest) (*pb.GeneralReply, error) {
	logrus.Printf("PostNewGame request：%s\n", in)
	collection := db.Database(setting.DatabaseSetting.DBName).Collection(setting.DatabaseSetting.CollectionName)
	newGame := pb.PostNewGameRequest{Game:in.Game}
	result, err := collection.InsertOne(context.TODO(), newGame)
	if err != nil {
		logrus.Fatal(err.Error())
		return &pb.GeneralReply{Result:err.Error()}, err
	}
	return &pb.GeneralReply{Result:fmt.Sprintf("Create a new game: %v",result.InsertedID)}, err
}

func mongoDB(port int) (){

	var err error

	for{
		logrus.Printf("Connecting to MongoDB...")
		mongoCtx, _ = context.WithTimeout(context.Background(), 5*time.Second)
		dbUrl := fmt.Sprintf("mongodb://%v:%v@%v:%v", setting.DatabaseSetting.Account, setting.DatabaseSetting.Password, setting.DatabaseSetting.ServerIP, port)
		db, err = mongo.Connect(mongoCtx, options.Client().ApplyURI(dbUrl))
		err = db.Ping(mongoCtx, nil)
		if err!=nil{
			logrus.Printf("Could not connect to MongoDB: %v\n", err.Error())
			break
		}
		logrus.Printf("Connected to Mongodb successfully")
		break
	}
}

//func test()  {
//	collection := db.Database("mydb").Collection("persons")
//	ruan := Person{"Ruan", 34, "Cape Town"}
//	_, err := collection.InsertOne(context.TODO(), ruan)
//	if err != nil {
//		log.Fatal(err)
//	}
//	doc2 := Person2{"Ruan", 34}
//	_, err = collection.InsertOne(context.TODO(), doc2)
//	if err != nil {
//		log.Fatal(err)
//	}
//}

func main() {

	// Init config
	setting.Setup()

	// Init database
	mongoDB(setting.DatabaseSetting.Port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", setting.ServerSetting.Port))
	if err != nil {
		logrus.Fatalf("Can't listen on port：%v", err.Error())
	}

	// Create a new gRPC server
	s := grpc.NewServer()
	pb.RegisterServiceServerServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("Can't init gRPC server：%v", err.Error())
	}

	defer mongoCtx.Done()
}