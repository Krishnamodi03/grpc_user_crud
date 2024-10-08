// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"go.mongodb.org/mongo-driver/mongo"
	"grpc/grpc_user_crud/controllers"
	"grpc/grpc_user_crud/repositories"
	"grpc/grpc_user_crud/services"
)

// Injectors from wire.go:

func InitializeUserCrudController(db *mongo.Collection) *controllers.UserCrudController {
	userRepositoryInterface := repositories.NewUserRepository(db)
	userServiceInterface := services.NewUserService(userRepositoryInterface)
	userCrudController := controllers.NewUserCrudController(userServiceInterface)
	return userCrudController
}
