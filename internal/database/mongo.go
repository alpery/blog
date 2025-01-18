package database

import (
	"blog/internal/model"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client *mongo.Client
	//TODO: add a real blog post here
	// Posts    []models.BlogPost
}

func InitMongoDB(uri string) (*MongoDB, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	return &MongoDB{Client: client}, nil
}

func (db *MongoDB) GetUserByEmail(email string) (*model.User, error) {
	collection := db.Client.Database("blog").Collection("users")

	filter := bson.D{{"email", email}}
	var user model.User
	if err := collection.FindOne(context.Background(), filter).Decode(&user); err != nil {
		return nil, fmt.Errorf("Failed to find user: %v", err)
	}

	return &user, nil
}

func (db *MongoDB) CreateBlog(post model.BlogPost) (interface{}, error) {
	collection := db.Client.Database("blog").Collection("posts")

	// Create a new blog post
	_, err := collection.InsertOne(context.Background(), post)
	if err != nil {
		return nil, fmt.Errorf("Failed to create blog post: %v", err)
	}

	return post, nil
}

func (db *MongoDB) GetBlogByID(id string) (interface{}, error) {
	collection := db.Client.Database("blog").Collection("posts")

	filter := bson.D{{"id", id}}
	var post model.BlogPost
	if err := collection.FindOne(context.Background(), filter).Decode(&post); err != nil {
		return nil, fmt.Errorf("Failed to find blog post: %v", err)
	}

	return post, nil
}

func (db *MongoDB) GetBlogByTitle(title string) (interface{}, error) {
	collection := db.Client.Database("blog").Collection("posts")

	filter := bson.D{{"title", title}}
	var post model.BlogPost
	if err := collection.FindOne(context.Background(), filter).Decode(&post); err != nil {
		return nil, fmt.Errorf("Failed to find blog post: %v", err)
	}

	return post, nil
}
