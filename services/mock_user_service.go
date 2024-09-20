package services

import (
	"context"
	"grpc/grpc_user_crud/models"

	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Create(ctx context.Context, name, email, phone, password string) (string, error) {
	args := m.Called(ctx, name, email, phone, password)
	return args.String(0), args.Error(1)
}

func (m *MockUserService) Read(ctx context.Context, id string) (*models.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserService) Update(ctx context.Context, id, name, email, phone, password string) error {
	args := m.Called(ctx, id, name, email, phone, password)
	return args.Error(0)
}

func (m *MockUserService) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
