package services

import (
	"context"
	"errors"
	"grpc/grpc_user_crud/models"
	"grpc/grpc_user_crud/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserService_Create(t *testing.T) {
	// Setup
	mockRepo := new(repositories.MockUserRepository)
	userService := UserService{Repo: mockRepo}

	user := &models.User{
		ID:       primitive.NewObjectID(),
		Name:     "Test User",
		Email:    "test@example.com",
		Phone:    "1234567890",
		Password: "password123",
	}

	mockRepo.On("Create",
		mock.Anything,
		mock.AnythingOfType("*models.User")).Return(user.ID.Hex(), nil)

	// Execute
	id, err := userService.Create(context.TODO(),
		user.Name,
		user.Email,
		user.Phone,
		user.Password)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, user.ID.Hex(), id)
	mockRepo.AssertExpectations(t)
}
func TestUserService_Create_Error(t *testing.T) {
	mockRepo := new(repositories.MockUserRepository)
	userService := UserService{Repo: mockRepo}

	mockRepo.On("Create",
		mock.Anything,
		mock.AnythingOfType("*models.User")).Return("",
		errors.New("failed to create user"))

	id, err := userService.Create(context.TODO(), "Test User", "test@example.com", "1234567890", "password123")

	assert.Error(t, err)
	assert.Equal(t, "", id)
	mockRepo.AssertExpectations(t)
}
func TestUserService_Read(t *testing.T) {
	mockRepo := new(repositories.MockUserRepository)
	userService := UserService{Repo: mockRepo}

	userID := primitive.NewObjectID().Hex()
	user := &models.User{
		ID:       primitive.NewObjectID(),
		Name:     "Test User",
		Email:    "test@example.com",
		Phone:    "1234567890",
		Password: "password123",
	}
	mockRepo.On("Read", mock.Anything, userID).Return(user, nil)
	// Execute
	result, err := userService.Read(context.TODO(), userID)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockRepo.AssertExpectations(t)
}
func TestUSerService_Read_Error(t *testing.T) {
	mockRepo := new(repositories.MockUserRepository)
	userService := UserService{Repo: mockRepo}
	mockRepo.On("Read", mock.Anything, "nonexistingId").Return((*models.User)(nil), errors.New("user not found"))

	result, err := userService.Read(context.TODO(), "nonexistingId")

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
func TestUserService_Update(t *testing.T) {
	mockRepo := new(repositories.MockUserRepository)
	userService := UserService{Repo: mockRepo}

	userID := primitive.NewObjectID()
	user := &models.User{
		ID:       userID,
		Name:     "Updated User",
		Email:    "updated@example.com",
		Phone:    "0987654321",
		Password: "newpassword123",
	}

	mockRepo.On("Update", mock.Anything,
		mock.AnythingOfType("*models.User")).Return(nil)

	// execute
	err := userService.Update(context.TODO(),
		user.ID.Hex(),
		user.Name,
		user.Email,
		user.Phone,
		user.Password)
	// Assert
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
func TestUserService_Update_Error(t *testing.T) {
	mockRepo := new(repositories.MockUserRepository)
	userService := UserService{Repo: mockRepo}

	userID := primitive.NewObjectID()

	mockRepo.On("Update", mock.Anything, mock.AnythingOfType("*models.User")).Return(errors.New("failed to update user"))
	err := userService.Update(context.TODO(), userID.Hex(), "Updated User", "updated@example.com", "0987654321", "newpassword123")

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}
func TestUserService_Delete(t *testing.T) {
	mockRepo := new(repositories.MockUserRepository)
	userService := UserService{Repo: mockRepo}
	userID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, userID).Return(nil)

	// Execute
	err := userService.Delete(context.TODO(), userID)
	// Assert
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
func TestUserService_Delete_Error(t *testing.T) {
	mockRepo := new(repositories.MockUserRepository)
	userServices := UserService{Repo: mockRepo}

	userID := primitive.NewObjectID().Hex()

	mockRepo.On("Delete", mock.Anything, userID).Return(errors.New("failed to delete user"))

	err := userServices.Delete(context.TODO(), userID)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}
