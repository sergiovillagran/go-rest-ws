package repository

import (
	"context"

	"github.com/sergiovillagran/rest-ws/models"
)

type Repository interface {
	InsertPost(ctx context.Context, post *models.Post) error
	GetUserById(ctx context.Context, Id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, Email string) (*models.User, error)
	InsertUser(ctx context.Context, user *models.User) error
	GetPostById(ctx context.Context, id string) (*models.Post, error)
	Close() error
}

var Implementation Repository

func SetRepository(repository Repository) {
	Implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return Implementation.InsertUser(ctx, user)
}

func GetUser(ctx context.Context, Id string) (*models.User, error) {
	return Implementation.GetUserById(ctx, Id)
}

func GetUserByEmail(ctx context.Context, Email string) (*models.User, error) {
	return Implementation.GetUserByEmail(ctx, Email)
}

func InsertPost(ctx context.Context, post *models.Post) error {
	return Implementation.InsertPost(ctx, post)
}

func GetPostById(ctx context.Context, id string) (*models.Post, error) {
	return Implementation.GetPostById(ctx, id)
}

func Close() error {
	return Implementation.Close()
}
