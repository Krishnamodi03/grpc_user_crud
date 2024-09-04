package controllers

import (
	"context"
	pb "grpc_user_crud/proto"

	"grpc_user_crud/services"
)

type UserCrudController struct {
	Service *services.UserService
	pb.UnimplementedUserCrudServiceServer
}

func (c *UserCrudController) CreateUser(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	id, err := c.Service.Create(ctx, req.GetName(), req.GetEmail(), req.GetPhone(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	return &pb.CreateResponse{Id: id}, nil
}

func (c *UserCrudController) ReadUser(ctx context.Context, req *pb.ReadRequest) (*pb.ReadResponse, error) {
	user, err := c.Service.Read(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.ReadResponse{Id: user.ID.Hex(), Name: user.Name, Email: user.Email, Phone: user.Phone, Password: user.Password}, nil
}

func (c *UserCrudController) UpdateUser(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	err := c.Service.Update(ctx, req.GetId(), req.GetName(), req.GetEmail(), req.GetPhone(), req.GetPassword())
	if err != nil {
		return nil, err
	}
	return &pb.UpdateResponse{Success: true}, nil
}
func (c *UserCrudController) DeleteUser(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := c.Service.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.DeleteResponse{Success: true}, nil
}
