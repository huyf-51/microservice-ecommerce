package grpc

import (
	"context"
	"log"
	"net"
	"payment/grpc/pb"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCreateOrderServiceServer
}

func (s *server) CreateOrderPayment(_ context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	log.Printf("order id: %v", in.GetOrderId())
	return &pb.CreateOrderResponse{Status: 1}, nil
}

func ConnectGrpcServer() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	pb.RegisterCreateOrderServiceServer(s, &server{})
	s.Serve(lis)
}
