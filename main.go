package main

import (
	"grpc/grpc_user_crud/database"
	"log"
	"net"

	pb "grpc/grpc_user_crud/proto"

	"grpc/grpc_user_crud/wire"

	"google.golang.org/grpc"
)

func main() {
	collection := database.DB
	// Setup repository, service, and controller

	// (old way which is manual dependency injection)
	// userRepo := &repositories.UserRepository{Collection: collection}
	// userService := &services.UserService{Repo: userRepo}
	// crudController := &controllers.UserCrudController{Service: userService}

	// (New way OF DEPENDENCY INJECTION using wire)
	crudController := wire.InitializeUserCrudController(collection)

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
