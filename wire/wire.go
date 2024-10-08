//go:build wireinject
// +build wireinject

// wire.go

package wire

import (
	"grpc/grpc_user_crud/controllers"
	"grpc/grpc_user_crud/repositories"
	"grpc/grpc_user_crud/services"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitializeUserCrudController(db *mongo.Collection) *controllers.UserCrudController {
	wire.Build(repositories.NewUserRepository, services.NewUserService, controllers.NewUserCrudController)
	return &controllers.UserCrudController{}
}
