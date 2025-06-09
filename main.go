package main

import (
	Config "GRPC/Config"
	Server "GRPC/Handler"
	"GRPC/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	Config.DbConnect()
	db := Config.DB

	lis, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterUserServiceServer(s, &Server.Server{DB: db})
	pb.RegisterFarmerServiceServer(s, &Server.FarmerServer{DB: db})
	pb.RegisterCategoryServiceServer(s, &Server.CategoryServer{})
	pb.RegisterCommodityServiceServer(s, &Server.CommodityServer{})
	pb.RegisterBidServiceServer(s, &Server.BidServer{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
