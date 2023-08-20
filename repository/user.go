package repository

import (
	"context"
	"vincentcore_interview/model/userModel"
)

func (r *repository) Create(ctx context.Context, user userModel.User) (userModel.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return userModel.User{}, err
	}
	return user, nil
}

func (r *repository) DetailByID(id string) (userModel.User, error) {
	var user userModel.User
	if err := r.db.Model(&userModel.User{}).Where("deleted_at IS NULL AND id = ?", id).First(&user).Error; err != nil {
		return userModel.User{}, err
	}

	return user, nil
}

func (r *repository) DetailByEmail(email string) (userModel.User, error) {
	var user userModel.User
	if err := r.db.Model(&userModel.User{}).Where("deleted_at IS NULL AND email = ?", email).First(&user).Error; err != nil {
		return userModel.User{}, err
	}

	return user, nil
}

func (r *repository) DetailByUsername(username string) (userModel.User, error) {
	var user userModel.User
	if err := r.db.Model(&userModel.User{}).Where("deleted_at IS NULL AND username = ?", username).First(&user).Error; err != nil {
		return userModel.User{}, err
	}

	return user, nil
}
