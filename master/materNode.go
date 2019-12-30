package main

import (
	"Simple_Distributed_System/pb"
	"Simple_Distributed_System/setting"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

type server struct{}

func (s *server) GetScore(ctx context.Context, in *pb.GetScoreRequest) (*pb.GetScoreReply, error) {
	logrus.Printf("GetScore request：%s\n", in.Game)
	return &pb.GetScoreReply{Home:[]int32{10,20,30},HomeTotal:60, Visitor:[]int32{30,40,50}, VisitorTotal:120}, nil
}

func (s *server) PutScore(ctx context.Context, in *pb.PutScoreRequest) (*pb.GeneralReply, error) {
	logrus.Printf("PutScore request：%s\n", in)
	return &pb.GeneralReply{Result:"ok"}, nil
}

func (s *server) GetGameList(ctx context.Context, in *pb.GeneralRequest) (*pb.GetGameListReply, error) {
	logrus.Printf("GetGameList request：%s\n", in)
	return &pb.GetGameListReply{Game:[]string{"Game1", "Game2"}}, nil
}

func (s *server) PostNewGame(ctx context.Context, in *pb.PostNewGameRequest) (*pb.GeneralReply, error) {
	logrus.Printf("PostNewGame request：%s\n", in)
	return &pb.GeneralReply{Result:"ok"}, nil
}

func mongoDB(port int) (*mongo.Client, context.Context, error){
	var db *mongo.Client
	var mongoCtx context.Context
	var rspErr error

	for{
		logrus.Printf("Connecting to MongoDB...")
		mongoCtx, _ = context.WithTimeout(context.Background(), 10*time.Second)
		db, rspErr = mongo.Connect(mongoCtx, options.Client().ApplyURI(fmt.Sprintf("mongodb://localhost:%v", port)))
		if rspErr!=nil{
			logrus.Printf("Could not connect to MongoDB: %v\n", rspErr)
			break
		}
		logrus.Printf("Connected to Mongodb successfully")
		break
	}
	return db, mongoCtx, rspErr
}

func main() {

	// Init config
	setting.Setup()

	// Init database
	_, mongoCtx, err := mongoDB(setting.DatabaseSetting.Port)
	if err != nil {
		logrus.Fatalf("Can't listen on port：%v", err.Error())
	}

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