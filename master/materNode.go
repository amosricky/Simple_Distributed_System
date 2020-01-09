package main

import (
	"Simple_Distributed_System/pb"
	"Simple_Distributed_System/setting"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

type GameItem struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Game 	string `bson:"game"`
	Score 	ScoreStruct `bson:"score"`
}

type ScoreStruct struct {
	Home [9]int32 `bson:"home"`
	Visitor [9]int32 `bson:"visitor"`
}

var db *mongo.Client
var mongoCtx context.Context
type server struct{}

func (s *server) GetScore(ctx context.Context, in *pb.GetScoreRequest) (*pb.GetScoreReply, error){
	logrus.Printf("GetScore request：%s\n", in.ID)
	var result GameItem

	for{
		queryCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		collection := db.Database(setting.DatabaseSetting.DBName).Collection(setting.DatabaseSetting.CollectionName)
		objectID, err := primitive.ObjectIDFromHex(in.ID)
		if err != nil {
			logrus.Fatal(err.Error())
			return &pb.GetScoreReply{}, err
		}

		err = collection.FindOne(queryCtx, bson.M{"_id": objectID}).Decode(&result)
		if err != nil {
			logrus.Fatal(err.Error())
			return &pb.GetScoreReply{}, err
		}
		break
	}

	//return &pb.GetScoreReply{Home:[]int32{10,20,30},HomeTotal:60, Visitor:[]int32{30,40,50}, VisitorTotal:120}, nil

	return &pb.GetScoreReply{Home:result.Score.Home[:], HomeTotal:100, Visitor:result.Score.Visitor[:], VisitorTotal:200}, nil
}

func (s *server) PutScore(ctx context.Context, in *pb.PutScoreRequest) (*pb.GeneralReply, error) {
	logrus.Printf("PutScore request：%s\n", in)

	//for{
	//	queryCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//	collection := db.Database(setting.DatabaseSetting.DBName).Collection(setting.DatabaseSetting.CollectionName)
	//	objectID, err := primitive.ObjectIDFromHex(in.ID)
	//	initScore := [9]int{}
	//	if err != nil {
	//		logrus.Fatal(err.Error())
	//		return &pb.GeneralReply{Result:err.Error()}, err
	//	}
	//}
	return &pb.GeneralReply{Result:"ok"}, nil
}

func (s *server) GetGameList(ctx context.Context, in *pb.GeneralRequest) (*pb.GetGameListReply, error) {
	logrus.Printf("GetGameList request：%s\n", in)
	var result []*pb.GameItem

	for{
		queryCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		collection := db.Database(setting.DatabaseSetting.DBName).Collection(setting.DatabaseSetting.CollectionName)
		cur, err := collection.Find(queryCtx, bson.M{})
		if err != nil {
			logrus.Fatal(err.Error())
			break
		}

		for cur.Next(context.TODO()) {
			data := &GameItem{}
			err := cur.Decode(data)
			if err != nil {
				logrus.Fatal(err.Error())
				break
			}
			result = append(result, &pb.GameItem{ID:data.ID.Hex(), Game:data.Game})
		}
		break
	}
	return &pb.GetGameListReply{Game:result}, nil
}

func (s *server) PostNewGame(ctx context.Context, in *pb.PostNewGameRequest) (*pb.GeneralReply, error) {
	logrus.Printf("PostNewGame request：%s\n", in)
	collection := db.Database(setting.DatabaseSetting.DBName).Collection(setting.DatabaseSetting.CollectionName)
	initScore := ScoreStruct{Home:[9]int32{}, Visitor:[9]int32{}}
	newGame := GameItem{Game:in.Game, Score:initScore,}
	result, err := collection.InsertOne(context.TODO(), newGame)
	if err != nil {
		logrus.Fatal(err.Error())
		return &pb.GeneralReply{Result:err.Error()}, err
	}
	return &pb.GeneralReply{Result:fmt.Sprintf("Create a new game: %v",result.InsertedID)}, nil
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