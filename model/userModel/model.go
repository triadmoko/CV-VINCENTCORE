package userModel

import "time"

type (
	User struct {
		UidUser   string
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time
		Username  string
		Email     string
		Password  string
	}
)

func (t *User) TableName() string {
	return "users"
}
