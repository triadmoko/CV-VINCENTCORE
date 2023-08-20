package service

import (
	"context"
	"vincentcore_interview/model/userModel"
	"vincentcore_interview/repository"

	"github.com/sirupsen/logrus"
)

type Service interface {
	Register(ctx context.Context, req userModel.RequestUser) (userModel.ResponseUser, int, error)
	Login(ctx context.Context, req userModel.RequestUserLogin) (userModel.ResponseUserLogin, int, error)
}
type service struct {
	repository repository.Repository
	logging    *logrus.Logger
}

func NewService(repository repository.Repository, logging *logrus.Logger) *service {
	return &service{repository: repository, logging: logging}
}
