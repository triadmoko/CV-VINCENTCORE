package service

import (
	"context"
	"net/http"
	"vincentcore_interview/model/userModel"
	"vincentcore_interview/pkg/constant"
	"vincentcore_interview/pkg/helpers"

	"gorm.io/gorm"
)

func (s *service) Register(ctx context.Context, req userModel.RequestUser) (userModel.ResponseUser, int, error) {
	resultUserByEmail, err := s.repository.DetailByEmail(req.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		s.logging.Error("s.repository.DetailByEmail ", err)
		return userModel.ResponseUser{}, http.StatusInternalServerError, constant.SomethingWentWrong
	}
	resultUserByUsername, err := s.repository.DetailByUsername(req.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		s.logging.Error("s.repository.DetailByUsername ", err)
		return userModel.ResponseUser{}, http.StatusInternalServerError, constant.SomethingWentWrong
	}

	if resultUserByEmail.UidUser != "" {
		return userModel.ResponseUser{}, http.StatusBadRequest, constant.UserHasBeenRegister
	}

	if resultUserByUsername.UidUser != "" {
		return userModel.ResponseUser{}, http.StatusBadRequest, constant.UserHasBeenRegister
	}

	password, err := helpers.HashPassword(req.Password)
	if err != nil {
		s.logging.Error("helpers.HashPassword", err.Error())
		return userModel.ResponseUser{}, http.StatusBadRequest, err
	}
	req.Password = password
	result, err := s.repository.Create(ctx, req.RequestFormatter())
	if err != nil {
		s.logging.Error("s.repository.Create", err.Error())
		return userModel.ResponseUser{}, http.StatusInternalServerError, constant.SomethingWentWrong
	}

	return result.ResponseFormatter(), http.StatusCreated, nil
}
