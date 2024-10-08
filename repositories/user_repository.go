package repositories

import (
	"context"
	"errors"
	"fmt"
	"grpc/grpc_user_crud/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrInvalidUserID = errors.New("invalid user ID")
	ErrCreateFailed  = errors.New("failed to create user")
	ErrUpdateFailed  = errors.New("failed to update user")
	ErrDeleteFailed  = errors.New("failed to delete user")
	ErrEmailPresent  = errors.New("email already exist")
	ErrPhonePresent  = errors.New("phone already exist")
)

type UserRepository struct {
	*mongo.Collection
}

type UserRepositoryInterface interface {
	Create(ctx context.Context, user *models.User) (string, error)
	Read(ctx context.Context, id string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id string) error
}

func NewUserRepository(collection *mongo.Collection) UserRepositoryInterface {
	return &UserRepository{Collection: collection}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) (string, error) {
	// check if already email is present
	if err := r.FindOne(ctx, bson.M{"email": user.Email}).Err(); err == nil {
		return "", ErrEmailPresent
	}
	// check if already phone is present
	if err := r.FindOne(ctx, bson.M{"phone": user.Phone}).Err(); err == nil {
		return "", ErrPhonePresent
	}
	result, err := r.InsertOne(ctx, user)
	if err != nil {
		return "", ErrCreateFailed
	}
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("failed to get inserted id")
	}
	return oid.Hex(), nil

}
func (r *UserRepository) Read(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidUserID
	}
	err = r.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to fetch user: %v", err)
	}
	return &user, nil
}
func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
	objID := user.ID
	updateObj := bson.M{
		"name":     user.Name,
		"email":    user.Email,
		"phone":    user.Phone,
		"password": user.Password,
	}

	result, err := r.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": updateObj})
	if err != nil {
		return ErrUpdateFailed
	}
	if result.MatchedCount == 0 {
		return ErrUserNotFound
	}
	return nil
}
func (r *UserRepository) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidUserID
	}
	result, err := r.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return ErrDeleteFailed
	}
	if result.DeletedCount == 0 {
		return ErrUserNotFound
	}
	return nil
}
