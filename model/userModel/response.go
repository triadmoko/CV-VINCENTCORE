package userModel

import "time"

type (
	ResponseUser struct {
		UidUser   string    `json:"uid_user"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Username  string    `json:"username"`
		Email     string    `json:"email"`
	}
	ResponseUserLogin struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Token    string `json:"token"`
	}
)
