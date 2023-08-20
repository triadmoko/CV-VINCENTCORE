package repository

import (
	"context"
	"vincentcore_interview/model/userModel"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, user userModel.User) (userModel.User, error)
	DetailByID(id string) (userModel.User, error)
	DetailByEmail(email string) (userModel.User, error)
	DetailByUsername(username string) (userModel.User, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}
