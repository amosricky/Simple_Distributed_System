package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"

	"../pb"
)

func main() {
	// 連線到遠端 gRPC 伺服器。
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("連線失敗：%v", err)
	}
	defer conn.Close()

	// 建立新的 Calculator 客戶端，所以等一下就能夠使用 Calculator 的所有方法。
	c := pb.NewGetScoreClient(conn)

	// 傳送新請求到遠端 gRPC 伺服器 Calculator 中，並呼叫 Plus 函式，讓兩個數字相加。
	result, err := c.GetScore(context.Background(), &pb.GetScoreRequest{Game:"testGame"})
	if err != nil {
		log.Fatalf("無法執行 GetScore 函式：%v", err)
	}else {
		resultJson, _ := json.Marshal(result)
		fmt.Printf("回傳結果：%s", resultJson)
	}


	putScoreClient := pb.NewPutScoreClient(conn)
	result, err = putScoreClient.PutScore(context.Background(), &pb.PutScoreRequest{Game:"testGame", Team:0, Round:2, Add:2})
}