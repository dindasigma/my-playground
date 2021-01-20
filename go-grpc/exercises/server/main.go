package main

import (
	"log"
	"net"

	"github.com/dindasigma/my-playground/go-grpc/exercises/server/controllers"
	pb "github.com/dindasigma/my-playground/go-grpc/exercises/server/proto/order"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	// seeder
	seeder()

	// create listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create grpc server
	s := grpc.NewServer()

	// attach controllers.Server to the order server
	pb.RegisterOrderManagementServer(s, &controllers.Server{})
	log.Printf("gRPC serve on %v", port)

	// serve
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func seeder() {
	controllers.OrderMap["102"] = pb.Order{Id: "102", Items: []string{"Google Pixel 3A", "Mac Book Pro"}, Destination: "Mountain View, CA", Price: 1800.00}
	controllers.OrderMap["103"] = pb.Order{Id: "103", Items: []string{"Apple Watch S4"}, Destination: "San Jose, CA", Price: 400.00}
	controllers.OrderMap["104"] = pb.Order{Id: "104", Items: []string{"Google Home Mini", "Google Nest Hub"}, Destination: "Mountain View, CA", Price: 400.00}
	controllers.OrderMap["105"] = pb.Order{Id: "105", Items: []string{"Amazon Echo"}, Destination: "San Jose, CA", Price: 30.00}
	controllers.OrderMap["106"] = pb.Order{Id: "106", Items: []string{"Amazon Echo", "Apple iPhone XS"}, Destination: "Mountain View, CA", Price: 300.00}
}