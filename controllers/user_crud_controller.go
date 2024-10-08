package controllers

import (
	"context"
	"errors"
	pb "grpc/grpc_user_crud/proto"
	"grpc/grpc_user_crud/repositories"
	"net/mail"
	"regexp"

	"grpc/grpc_user_crud/services"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserCrudController struct {
	Service services.UserServiceInterface
	pb.UnimplementedUserCrudServiceServer
}

func NewUserCrudController(service services.UserServiceInterface) *UserCrudController {
	return &UserCrudController{Service: service}
}

func (c *UserCrudController) CreateUser(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	name := req.GetName()
	if len(name) < 2 || len(name) > 30 || !regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(name) {
		return nil, status.Errorf(codes.InvalidArgument, "Name must be between 2 and 50 characters and contain only alphabetic characters")
	}

	email := req.GetEmail()
	_, errMail := mail.ParseAddress(email)
	if email == "" || errMail != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Email Address")
	}

	phone := req.GetPhone()
	if len(phone) != 10 || !regexp.MustCompile(`^[0-9]+$`).MatchString(phone) {
		return nil, status.Errorf(codes.InvalidArgument, "Phone number must be exactly 10 digits and numeric")
	}

	password := req.GetPassword()
	if len(password) < 8 {
		return nil, status.Errorf(codes.InvalidArgument, "Password must be atleast 8 characters long")
	}
	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return nil, status.Errorf(codes.InvalidArgument, "Password must contain at least one uppercase letter")
	}
	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return nil, status.Errorf(codes.InvalidArgument, "Password must contain at least one lowercase letter")
	}
	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return nil, status.Errorf(codes.InvalidArgument, "Password must contain at least one digit")
	}
	if !regexp.MustCompile(`[!@#\$%\^&\*]`).MatchString(password) {
		return nil, status.Errorf(codes.InvalidArgument, "Password must contain at least one special character (!, @, #, $, etc.)")
	}
	id, err := c.Service.Create(ctx, name, email, phone, password)
	if err != nil {
		if errors.Is(err, repositories.ErrEmailPresent) {
			return nil, status.Errorf(codes.AlreadyExists, "email already exist")
		}
		if errors.Is(err, repositories.ErrPhonePresent) {
			return nil, status.Errorf(codes.AlreadyExists, "phone already exist")
		}
		return nil, status.Errorf(codes.Internal, "Failed to create user : %v", err)
	}
	return &pb.CreateResponse{Id: id}, nil
}

func (c *UserCrudController) ReadUser(ctx context.Context, req *pb.ReadRequest) (*pb.ReadResponse, error) {
	id := req.GetId()
	if id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "ID is required")
	}
	if !primitive.IsValidObjectID(id) {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid ID format")
	}
	user, err := c.Service.Read(ctx, id)
	if err != nil {
		if errors.Is(err, repositories.ErrUserNotFound) {
			return nil, status.Errorf(codes.NotFound, "User with ID %s not found", id)
		}
		return nil, status.Errorf(codes.Internal, "Failed to read user:%v", err)
	}
	return &pb.ReadResponse{Id: user.ID.Hex(),
		Name:     user.Name,
		Email:    user.Email,
		Phone:    user.Phone,
		Password: user.Password,
	}, nil
}

func (c *UserCrudController) UpdateUser(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	id := req.GetId()
	if id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "ID is required")
	}
	if !primitive.IsValidObjectID(id) {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid ID format")
	}
	name := req.GetName()
	if len(name) < 2 || len(name) > 30 || !regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(name) {
		return nil, status.Errorf(codes.InvalidArgument, "Name must be between 2 and 50 characters and contain only alphabetic characters")
	}

	email := req.GetEmail()
	_, errMail := mail.ParseAddress(email)
	if email == "" || errMail != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Email Address")
	}

	phone := req.GetPhone()
	if len(phone) != 10 || !regexp.MustCompile(`^[0-9]+$`).MatchString(phone) {
		return nil, status.Errorf(codes.InvalidArgument, "Phone number must be exactly 10 digits and numeric")
	}

	password := req.GetPassword()
	if len(password) < 8 {
		return nil, status.Errorf(codes.InvalidArgument, "Password must be atleast 8 characters long")
	}
	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return nil, status.Errorf(codes.InvalidArgument, "Password must contain at least one uppercase letter")
	}
	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return nil, status.Errorf(codes.InvalidArgument, "Password must contain at least one lowercase letter")
	}
	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return nil, status.Errorf(codes.InvalidArgument, "Password must contain at least one digit")
	}
	if !regexp.MustCompile(`[!@#\$%\^&\*]`).MatchString(password) {
		return nil, status.Errorf(codes.InvalidArgument, "Password must contain at least one special character (!, @, #, $, etc.)")
	}

	err := c.Service.Update(ctx, id, name, email, phone, password)

	if err != nil {
		if errors.Is(err, repositories.ErrUserNotFound) {
			return nil, status.Errorf(codes.NotFound, "User with ID %s not found", id)
		}
		return nil, status.Errorf(codes.Internal, "Failed to update user: %v", err)
	}

	return &pb.UpdateResponse{Success: true}, nil
}

func (c *UserCrudController) DeleteUser(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	id := req.GetId()
	if id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "ID is required")
	}
	if !primitive.IsValidObjectID(id) {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid ID format")
	}
	err := c.Service.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, repositories.ErrUserNotFound) {
			return nil, status.Errorf(codes.NotFound, "User with ID %s not found", id)
		}
		return nil, status.Errorf(codes.Internal, "Failed to delete user : %v", err)
	}
	return &pb.DeleteResponse{Success: true}, nil
}
