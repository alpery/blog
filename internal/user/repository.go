package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(ctx context.Context, user *BlogUser) error
	FindByUsername(ctx context.Context, username string) (*BlogUser, error)
}

type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, user *BlogUser) error {
	collection := r.db.Collection("users")
	_, err := collection.InsertOne(ctx, user)
	return err
}

func (r *repository) FindByUsername(ctx context.Context, username string) (*BlogUser, error) {
	collection := r.db.Collection("users")
	var user BlogUser
	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
