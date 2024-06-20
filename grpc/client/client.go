package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	pb "gofortest/grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewMeatServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetMeatAmounts(ctx, &pb.MeatRequest{MeatList: "Fatback t-bone t-bone, pastrami  ..   t-bone.  pork, meatloaf jowl enim.  Bresaola t-bone."})
	if err != nil {
		log.Fatalf("could not get meat amounts: %v", err)
	}

	// Convert the response to JSON
	jsonResponse, err := json.Marshal(r)
	if err != nil {
		log.Fatalf("failed to marshal response to JSON: %v", err)
	}

	log.Printf("%s", jsonResponse)
}
