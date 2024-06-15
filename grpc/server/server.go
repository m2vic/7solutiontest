package main

import (
	"context"
	"log"
	"net"
	"regexp"
	"strings"

	pb "gofortest/grpc/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMeatServiceServer
}

func (server) GetMeatAmounts(ctx context.Context, req *pb.MeatRequest) (*pb.MeatResponse, error) {
	meatList := req.GetMeatList()

	re := regexp.MustCompile(`[a-zA-Z-]+`)
	words := re.FindAllString(strings.ToLower(meatList), -1)

	meatCounts := make(map[string]int32)
	for _, word := range words {
		meatCounts[word]++
	}

	return &pb.MeatResponse{MeatItems: meatCounts}, nil

}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterMeatServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
