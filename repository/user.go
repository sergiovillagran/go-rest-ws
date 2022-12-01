package repository

import (
	"context"

	"github.com/sergiovillagran/rest-ws/models"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, Id int64) (*models.User, error)
}

var Implementation UserRepository

func SetRepository(repository UserRepository) {
	Implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return Implementation.InsertUser(ctx, user)
}

func GetUser(ctx context.Context, Id int64) (*models.User, error) {
	return Implementation.GetUserById(ctx, Id)
}