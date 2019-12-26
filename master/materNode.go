package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"../pb"
)

// master 建構體會實作 Calculator 的 gRPC 伺服器。
type server struct{}

func (s *server) GetScore(ctx context.Context, in *pb.GetScoreRequest) (*pb.GetScoreReply, error) {
	fmt.Printf("GetScore request：%s", in.Game)
	// 包裝成 Protobuf 建構體並回傳。
	return &pb.GetScoreReply{Home:[]int32{10,20,30},HomeTotal:60, Visitor:[]int32{30,40,50}, VisitorTotal:120}, nil
}

func (s *server) PutScore(ctx context.Context, in *pb.PutScoreRequest) (*pb.GeneralReply, error) {
	fmt.Printf("PutScore request：%s", in)
	// 包裝成 Protobuf 建構體並回傳。
	return &pb.GeneralReply{Result:"ok"}, nil
}

func main() {
	// 監聽指定埠口，這樣服務才能在該埠口執行。
	fmt.Println("Start!")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Println("無法監聽該埠口：%v", err)
	}

	// 建立新 gRPC 伺服器並註冊 Calculator 服務。
	s := grpc.NewServer()
	pb.RegisterGetScoreServer(s, &server{})
	pb.RegisterPutScoreServer(s, &server{})

	// 在 gRPC 伺服器上註冊反射服務。
	reflection.Register(s)

	// 開始在指定埠口中服務。
	if err := s.Serve(lis); err != nil {
		log.Fatalf("無法提供服務：%v", err)
	}
}