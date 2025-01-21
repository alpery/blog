package blog

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Find(ctx context.Context, filter interface{}) ([]BlogPost, error)
	FindOne(ctx context.Context, id string) (*BlogPost, error)
	Create(ctx context.Context, post *BlogPost) error
}

type mongoRepository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) Repository {
	return &mongoRepository{db: db}
}

func (r *mongoRepository) Find(ctx context.Context, filter interface{}) ([]BlogPost, error) {
	collection := r.db.Collection("blogs")
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var posts []BlogPost
	if err = cursor.All(ctx, &posts); err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *mongoRepository) FindOne(ctx context.Context, id string) (*BlogPost, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	collection := r.db.Collection("blogs")
	var post BlogPost

	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&post)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *mongoRepository) Create(ctx context.Context, post *BlogPost) error {
	collection := r.db.Collection("blogs")
	post.CreatedAt = time.Now()
	_, err := collection.InsertOne(ctx, post)
	return err
}
