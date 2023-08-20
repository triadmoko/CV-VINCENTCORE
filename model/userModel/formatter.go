package userModel

import (
	"time"
	"vincentcore_interview/pkg/helpers"
)

func (req RequestUser) RequestFormatter() User {
	return User{
		UidUser:   helpers.UUID(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		DeletedAt: nil,
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password,
	}
}
func (req User) ResponseFormatter() ResponseUser {
	return ResponseUser{
		UidUser:   req.UidUser,
		CreatedAt: req.CreatedAt,
		UpdatedAt: req.UpdatedAt,
		Username:  req.Username,
		Email:     req.Email,
	}
}
