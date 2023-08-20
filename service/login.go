package service

import (
	"context"
	"net/http"
	"vincentcore_interview/model/userModel"
	"vincentcore_interview/pkg/constant"
	"vincentcore_interview/pkg/helpers"

	"gorm.io/gorm"
)

func (s *service) Login(ctx context.Context, req userModel.RequestUserLogin) (userModel.ResponseUserLogin, int, error) {
	user, err := s.repository.DetailByEmail(req.Email)
	if err == gorm.ErrRecordNotFound {
		return userModel.ResponseUserLogin{}, http.StatusNotFound, constant.UserNotFound
	}
	if err != nil {
		s.logging.Error("s.repository.DetailByEmail ", err)
		return userModel.ResponseUserLogin{}, http.StatusInternalServerError, constant.SomethingWentWrong
	}
	if err := helpers.ComparePassword(user.Password, req.Password); err != nil {
		return userModel.ResponseUserLogin{}, http.StatusUnauthorized, constant.PasswordOrEmailNotMatch
	}

	token, err := helpers.Sign(map[string]interface{}{
		"uid_user": user.UidUser,
		"email":    user.Email,
	}, 0)
	if err != nil {
		s.logging.Error("helpers.Sign ", err)
		return userModel.ResponseUserLogin{}, http.StatusUnauthorized, constant.SomethingWentWrong
	}
	return userModel.ResponseUserLogin{
			Username: user.Username,
			Email:    user.Email,
			Token:    token,
		},
		http.StatusOK, nil
}
