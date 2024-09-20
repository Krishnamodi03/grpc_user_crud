package controllers

import (
	"context"
	"errors"
	"grpc/grpc_user_crud/models"
	"grpc/grpc_user_crud/proto"
	"grpc/grpc_user_crud/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserController_Create(t *testing.T) {
	mockServices := new(services.MockUserService)
	controller := UserCrudController{Service: mockServices}

	req := &proto.CreateRequest{
		Name:     "Test User",
		Email:    "test@example.com",
		Phone:    "1234567890",
		Password: "password123",
	}
	expectedID := "someobjectID"
	mockServices.On("Create", context.Background(),
		req.GetName(),
		req.GetEmail(),
		req.GetPhone(),
		req.GetPassword()).Return(expectedID, nil)

	resp, err := controller.CreateUser(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedID, resp.GetId())
	mockServices.AssertExpectations(t)
}
func TestUserController_Create_Error(t *testing.T) {
	mockServices := new(services.MockUserService)
	controller := UserCrudController{Service: mockServices}

	req := &proto.CreateRequest{
		Name:     "Test User",
		Email:    "invalid-email",
		Phone:    "1234567890",
		Password: "password123",
	}

	mockServices.On("Create", context.Background(),
		req.GetName(),
		req.GetEmail(),
		req.GetPhone(),
		req.GetPassword()).Return("", errors.New("Invalid email format"))
	resp, err := controller.CreateUser(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	mockServices.AssertExpectations(t)
}
func TestUserController_Read(t *testing.T) {
	mockServices := new(services.MockUserService)
	controller := UserCrudController{Service: mockServices}
	userID := "someObjectID"
	user := &models.User{
		ID:       primitive.NewObjectID(),
		Name:     "Test User",
		Email:    "test@example.com",
		Phone:    "1234567890",
		Password: "password123",
	}
	req := &proto.ReadRequest{Id: userID}
	mockServices.On("Read", context.Background(), userID).Return(user, nil)
	resp, err := controller.ReadUser(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, user.ID.Hex(), resp.GetId())
	assert.Equal(t, user.Name, resp.Name)
	assert.Equal(t, user.Email, resp.Email)
	assert.Equal(t, user.Phone, resp.Phone)
	assert.Equal(t, user.Password, resp.Password)
	mockServices.AssertExpectations(t)
}
func TestUserController_Read_Error(t *testing.T) {
	mockServices := new(services.MockUserService)
	controller := UserCrudController{Service: mockServices}

	userID := "someObjectID"

	req := &proto.ReadRequest{Id: userID}

	mockServices.On("Read",
		context.Background(),
		userID).Return((*models.User)(nil), errors.New("User not found"))

	resp, err := controller.ReadUser(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	mockServices.AssertExpectations(t)
}
func TestUserController_Update(t *testing.T) {
	mockServices := new(services.MockUserService)
	controller := UserCrudController{Service: mockServices}
	userID := "someObjectID"
	req := &proto.UpdateRequest{
		Id:       userID,
		Name:     "Updated User",
		Email:    "updated@example.com",
		Phone:    "9876543210",
		Password: "updatedPassword",
	}
	mockServices.On("Update", context.Background(),
		userID,
		req.GetName(),
		req.GetEmail(),
		req.GetPhone(),
		req.GetPassword()).Return(nil)

	resp, err := controller.UpdateUser(context.Background(), req)
	assert.NoError(t, err)
	assert.True(t, resp.GetSuccess())
	mockServices.AssertExpectations(t)
}
func TestUserController_Update_Error(t *testing.T) {
	mockServices := new(services.MockUserService)
	controller := UserCrudController{Service: mockServices}

	userID := "invalidObjectId"

	req := &proto.UpdateRequest{
		Id:       userID,
		Name:     "Updated User",
		Email:    "updated@example.com",
		Phone:    "9876543210",
		Password: "updatedPassword",
	}
	mockServices.On("Update", context.Background(),
		userID,
		req.GetName(),
		req.GetEmail(),
		req.GetPhone(),
		req.GetPassword()).Return(errors.New("user not found"))
	resp, err := controller.UpdateUser(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	mockServices.AssertExpectations(t)
}
func TestUserController_Delete(t *testing.T) {
	mockServices := new(services.MockUserService)
	controller := UserCrudController{Service: mockServices}
	userID := "someObjectID"

	req := &proto.DeleteRequest{
		Id: userID,
	}
	mockServices.On("Delete", context.Background(), userID).Return(nil)

	resp, err := controller.DeleteUser(context.Background(), req)
	assert.NoError(t, err)
	assert.True(t, resp.GetSuccess())
	mockServices.AssertExpectations(t)
}
func TestUserController_Delete_Error(t *testing.T) {
	mockServices := new(services.MockUserService)
	controller := UserCrudController{Service: mockServices}

	userID := "invalidObjectID"

	req := &proto.DeleteRequest{
		Id: userID,
	}

	mockServices.On("Delete", context.Background(), userID).Return(errors.New("user not found"))

	resp, err := controller.DeleteUser(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	mockServices.AssertExpectations(t)
}
