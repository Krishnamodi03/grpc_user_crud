package services

import (
	"context"
	"errors"
	"grpc/grpc_user_crud/models"
	"grpc/grpc_user_crud/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	Repo repositories.UserRepositoryInterface
}

type UserServiceInterface interface {
	Create(ctx context.Context, name, email, phone, password string) (string, error)
	Read(ctx context.Context, id string) (*models.User, error)
	Update(ctx context.Context, id, name, email, phone, password string) error
	Delete(ctx context.Context, id string) error
}

func NewUserService(repo repositories.UserRepositoryInterface) UserServiceInterface {
	return &UserService{Repo: repo}
}

func (s *UserService) Create(ctx context.Context, name, email, phone, password string) (string, error) {
	user := &models.User{
		ID:       primitive.NewObjectID(),
		Name:     name,
		Email:    email,
		Phone:    phone,
		Password: password,
	}
	id, err := s.Repo.Create(ctx, user)
	if err != nil {
		if errors.Is(err, repositories.ErrEmailPresent) {
			return "", repositories.ErrEmailPresent
		}
		if errors.Is(err, repositories.ErrPhonePresent) {
			return "", repositories.ErrPhonePresent
		}
		return "", repositories.ErrCreateFailed
	}
	return id, nil
}
func (s *UserService) Read(ctx context.Context, id string) (*models.User, error) {
	user, err := s.Repo.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, repositories.ErrUserNotFound
	}
	return user, nil
}
func (s *UserService) Update(ctx context.Context, id, name, email, phone, password string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return repositories.ErrInvalidUserID
	}
	user := &models.User{
		ID:       oid,
		Name:     name,
		Email:    email,
		Phone:    phone,
		Password: password,
	}
	err = s.Repo.Update(ctx, user)
	if errors.Is(err, repositories.ErrUserNotFound) {
		return repositories.ErrUserNotFound
	}
	if err != nil {
		return repositories.ErrUpdateFailed
	}
	return nil
}
func (s *UserService) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return repositories.ErrInvalidUserID
	}
	err = s.Repo.Delete(ctx, oid.Hex())
	if errors.Is(err, repositories.ErrUserNotFound) {
		return repositories.ErrUserNotFound
	}
	if err != nil {
		return repositories.ErrDeleteFailed
	}
	return nil
}
