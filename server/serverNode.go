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

type server struct{}

func (s *server) GetScore(ctx context.Context, in *pb.GetScoreRequest) (*pb.GetScoreReply, error){
	logrus.Printf("GetScore request：%s\n", in.ID)
	var result GameItem

	for{
		db, dbErr := mongoDB(in.DbIP, int(in.DbPort))
		if dbErr != nil{
			logrus.Printf("Could not connect to MongoDB: %v\n", dbErr.Error())
			return &pb.GetScoreReply{}, dbErr
		}
		queryCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		defer queryCtx.Done()
		collection := db.Database(setting.DatabaseSetting.DBName).Collection(setting.DatabaseSetting.CollectionName)
		objectID, objectIDErr := primitive.ObjectIDFromHex(in.ID)
		if objectIDErr != nil {
			logrus.Warnf(objectIDErr.Error())
			return &pb.GetScoreReply{}, objectIDErr
		}

		findOneErr := collection.FindOne(queryCtx, bson.M{"_id": objectID}).Decode(&result)
		if findOneErr != nil {
			logrus.Warnf(findOneErr.Error())
			return &pb.GetScoreReply{}, findOneErr
		}
		break
	}

	homeScore := result.Score.Home[:]
	visitorScore := result.Score.Visitor[:]
	countHomeScore := int32(0)
	countVisitorScore := int32(0)

	for i:=0;i<9;i++{
		countHomeScore += homeScore[i]
		countVisitorScore += visitorScore[i]
	}

	return &pb.GetScoreReply{Home:homeScore, HomeTotal:countHomeScore, Visitor:visitorScore, VisitorTotal:countVisitorScore}, nil
}

func (s *server) PutScore(ctx context.Context, in *pb.PutScoreRequest) (*pb.GeneralReply, error) {
	logrus.Printf("PutScore request：%s\n", in)
	var getItem GameItem

	for{
		db, dbErr := mongoDB(setting.DatabaseSetting.ServerIP,setting.DatabaseSetting.Port)
		if dbErr != nil{
			logrus.Printf("Could not connect to MongoDB: %v\n", dbErr.Error())
			return &pb.GeneralReply{Result:dbErr.Error()}, dbErr
		}
		queryCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		defer queryCtx.Done()
		collection := db.Database(setting.DatabaseSetting.DBName).Collection(setting.DatabaseSetting.CollectionName)
		objectID, objectIDErr := primitive.ObjectIDFromHex(in.ID)
		if objectIDErr != nil {
			logrus.Warnf(objectIDErr.Error())
			return &pb.GeneralReply{Result:objectIDErr.Error()}, objectIDErr
		}

		FindOneErr := collection.FindOne(queryCtx, bson.M{"_id": objectID}).Decode(&getItem)
		if FindOneErr != nil {
			logrus.Warnf(FindOneErr.Error())
			return &pb.GeneralReply{Result:FindOneErr.Error()}, FindOneErr
		}

		if (in.Round > 9) || (in.Round < 1){
			logrus.Warnf("Round range must in 1~9")
			return &pb.GeneralReply{Result:"Round range must in 1~9"}, nil
		}

		switch in.Team.String() {
		case "Home":
			score := getItem.Score.Home
			score[in.Round-1] += in.Add

			update := bson.M{
				"$set": bson.M{"score.home":score},
			}

			// Result not use
			collection.FindOneAndUpdate(queryCtx, bson.M{"_id": objectID}, update)
		case "Visitor":
			score := getItem.Score.Visitor
			score[in.Round-1] += in.Add

			update := bson.M{
				"$set": bson.M{"score.visitor":score},
			}

			// Result not use
			collection.FindOneAndUpdate(queryCtx, bson.M{"_id": objectID}, update)
		default:
			return &pb.GeneralReply{Result:"Team not exist."}, nil
		}
		break
	}
	return &pb.GeneralReply{Result:"ok"}, nil
}

func (s *server) GetGameList(ctx context.Context, in *pb.GeneralRequest) (*pb.GetGameListReply, error) {
	logrus.Printf("GetGameList request：%s\n", in)
	var result []*pb.GameItem

	for{
		fmt.Printf("DBIP : %v \n", in.DbIP)
		fmt.Printf("DbPort : %v \n", in.DbPort)
		db, dbErr := mongoDB(in.DbIP, int(in.DbPort))
		if dbErr != nil{
			logrus.Printf("Could not connect to MongoDB: %v\n", dbErr.Error())
			return &pb.GetGameListReply{}, dbErr
		}
		queryCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		defer queryCtx.Done()
		collection := db.Database(setting.DatabaseSetting.DBName).Collection(setting.DatabaseSetting.CollectionName)
		cur, curErr := collection.Find(queryCtx, bson.M{})
		if curErr != nil {
			logrus.Warnf(curErr.Error())
			return &pb.GetGameListReply{}, curErr
		}

		for cur.Next(context.TODO()) {
			data := &GameItem{}
			decodeErr := cur.Decode(data)
			if decodeErr != nil {
				logrus.Warnf(decodeErr.Error())
				return &pb.GetGameListReply{}, decodeErr
			}
			result = append(result, &pb.GameItem{ID:data.ID.Hex(), Game:data.Game})
		}
		break
	}
	return &pb.GetGameListReply{Game:result}, nil
}

func (s *server) PostNewGame(ctx context.Context, in *pb.PostNewGameRequest) (*pb.GeneralReply, error) {
	logrus.Printf("PostNewGame request：%s\n", in)

	db, dbErr := mongoDB(setting.DatabaseSetting.ServerIP,setting.DatabaseSetting.Port)
	if dbErr != nil{
		logrus.Printf("Could not connect to MongoDB: %v\n", dbErr.Error())
		return &pb.GeneralReply{Result:dbErr.Error()}, dbErr
	}
	queryCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	defer queryCtx.Done()
	collection := db.Database(setting.DatabaseSetting.DBName).Collection(setting.DatabaseSetting.CollectionName)
	initScore := ScoreStruct{Home:[9]int32{}, Visitor:[9]int32{}}
	newGame := GameItem{Game:in.Game, Score:initScore,}
	result, insertOneErr := collection.InsertOne(queryCtx, newGame)
	if insertOneErr != nil {
		logrus.Fatal(insertOneErr.Error())
		return &pb.GeneralReply{Result:insertOneErr.Error()}, insertOneErr
	}
	return &pb.GeneralReply{Result:fmt.Sprintf("Create a new game: %v",result.InsertedID)}, nil
}

func mongoDB(ip string , port int) (*mongo.Client, error) {
	var db *mongo.Client
	var mongoCtx context.Context
	var err error

	for{
		logrus.Printf("Connecting to MongoDB...")
		mongoCtx, _ = context.WithTimeout(context.Background(), 5*time.Second)
		dbUrl := fmt.Sprintf("mongodb://%v:%v/?connect=direct", ip, port)
		fmt.Printf(dbUrl)
		db, err = mongo.Connect(mongoCtx, options.Client().ApplyURI(dbUrl))
		if err != nil{
			break
		}

		for times:=0; times<3;times++{
			err = db.Ping(mongoCtx, nil)
			if err!=nil{
				logrus.Printf("Could not connect to MongoDB: %v\n", err.Error())
				time.Sleep(3 * time.Second)
			}else {
				logrus.Printf("Connected to Mongodb successfully")
				err = nil
				break
			}
		}
		break
	}
	return db, err
}

func main() {
	logrus.Infof("Mater node start")

	// Init config
	setting.Setup()

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
}