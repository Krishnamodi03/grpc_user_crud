package repositories

import (
	"context"
	"grpc_user_crud/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) (string, error) {
	result, err := r.Collection.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}
	oid, _ := result.InsertedID.(primitive.ObjectID)
	return oid.Hex(), nil

}
func (r *UserRepository) Read(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = r.Collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, err
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

	_, err := r.Collection.UpdateOne(
		ctx,
		bson.M{"_id": objID},
		bson.M{"$set": updateObj})
	return err
}
func (r *UserRepository) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
