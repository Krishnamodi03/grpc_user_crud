package services

import (
	"context"
	"grpc_user_crud/models"
	"grpc_user_crud/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func (s *UserService) Create(ctx context.Context, name, email, phone, password string) (string, error) {
	user := &models.User{
		ID:       primitive.NewObjectID(),
		Name:     name,
		Email:    email,
		Phone:    phone,
		Password: password,
	}
	return s.Repo.Create(ctx, user)
}
func (s *UserService) Read(ctx context.Context, id string) (*models.User, error) {
	return s.Repo.Read(ctx, id)
}
func (s *UserService) Update(ctx context.Context, id, name, email, phone, password string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	user := &models.User{
		ID:       oid,
		Name:     name,
		Email:    email,
		Phone:    phone,
		Password: password,
	}
	return s.Repo.Update(ctx, user)
}
func (s *UserService) Delete(ctx context.Context, id string) error {
	return s.Repo.Delete(ctx, id)
}
