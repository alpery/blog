package blog

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type Service interface {
	GetPosts(ctx context.Context) ([]BlogPost, error)
	GetPost(ctx context.Context, id string) (*BlogPost, error)
	CreatePost(ctx context.Context, post *BlogPost) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetPosts(ctx context.Context) ([]BlogPost, error) {
	return s.repo.Find(ctx, bson.M{})
}

func (s *service) GetPost(ctx context.Context, id string) (*BlogPost, error) {
	return s.repo.FindOne(ctx, id)
}

func (s *service) CreatePost(ctx context.Context, post *BlogPost) error {
	return s.repo.Create(ctx, post)
}
