package main

import (
	"Simple_Distributed_System/pb"
	"Simple_Distributed_System/setting"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)


func main() {

	setting.Setup()
	serverUrl := fmt.Sprintf("localhost:%v", setting.ServerSetting.Port)
	conn, err := grpc.Dial(serverUrl, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("Can't connect to server：%v", err)
	}
	defer conn.Close()

	connGetScore := pb.NewServiceServerClient(conn)
	resultGetScore, err := connGetScore.GetScore(context.Background(), &pb.GetScoreRequest{ID:"5e16fa4d464087a28bac9e8b"})
	if err != nil {
		logrus.Fatalf("Can't execute [GetScore] function：%v", err)
	}else {
		resultJson, _ := json.Marshal(resultGetScore)
		logrus.Printf("Reply [GetScore]：%s", resultJson)
	}

	//connPutScore := pb.NewServiceServerClient(conn)
	//resultPutScore, err := connPutScore.PutScore(context.Background(), &pb.PutScoreRequest{ID:"5e16fa4d464087a28bac9e8b", Team:2, Round:2, Add:2})
	//if err != nil {
	//	logrus.Printf("Can't execute [PutScore] function：%v", err)
	//}else {
	//	resultJson, _ := json.Marshal(resultPutScore)
	//	logrus.Printf("Reply [PutScore]：%s", resultJson)
	//}

	//connGetGameList := pb.NewServiceServerClient(conn)
	//resultGetGameList, err := connGetGameList.GetGameList(context.Background(), &pb.GeneralRequest{})
	//if err != nil {
	//	logrus.Printf("Can't execute [GetGameList] function：%v", err)
	//}else {
	//	resultJson, _ := json.Marshal(resultGetGameList)
	//	logrus.Printf("Reply [GetGameList]：%s", resultJson)
	//}

	//connPostNewGame := pb.NewServiceServerClient(conn)
	//resultPostNewGame, err := connPostNewGame.PostNewGame(context.Background(), &pb.PostNewGameRequest{Game:"Test0109_3"})
	//if err != nil {
	//	logrus.Printf("Can't execute [PostNewGame] function：%v", err)
	//}else {
	//	resultJson, _ := json.Marshal(resultPostNewGame)
	//	logrus.Printf("Reply [PostNewGame]：%s", resultJson)
	//}
}