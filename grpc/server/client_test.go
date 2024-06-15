package main

import (
	"context"
	"log"
	"testing"

	pb "gofortest/grpc/proto"

	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {

	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterMeatServiceServer(s, &server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestGetMeatAmounts(t *testing.T) {
	ctx := context.Background()
	//resolver.SetDefaultScheme("passthrough")

	conn, err := grpc.NewClient("passthrough://bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewMeatServiceClient(conn)

	req := &pb.MeatRequest{MeatList: "Fatback t-bone t-bone, pastrami  ..   t-bone.  pork, meatloaf jowl enim.  Bresaola t-bone."}
	resp, err := client.GetMeatAmounts(ctx, req)
	if err != nil {
		t.Fatalf("GetMeatAmounts failed: %v", err)
	}

	expected := map[string]int32{
		"fatback":  1,
		"t-bone":   4,
		"pastrami": 1,
		"pork":     1,
		"meatloaf": 1,
		"jowl":     1,
		"enim":     1,
		"bresaola": 1,
	}

	for name, amount := range expected {
		if resp.MeatItems[name] != amount {
			t.Errorf("expected %v: %v, got %v", name, amount, resp.MeatItems[name])
		}
	}
}
