package main

import (
	"grpc_user_crud/controllers"
	"grpc_user_crud/database"
	"grpc_user_crud/repositories"
	"grpc_user_crud/services"
	"log"
	"net"

	pb "grpc_user_crud/proto"

	"google.golang.org/grpc"
)

func main() {
	collection := database.DB
	// Setup repository, service, and controller
	userRepo := &repositories.UserRepository{Collection: collection}
	userService := &services.UserService{Repo: userRepo}
	crudController := &controllers.UserCrudController{Service: userService}

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserCrudServiceServer(grpcServer, crudController)

	log.Printf("server listening at %v", lis.Addr())

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve : %v", err)
	}
}
