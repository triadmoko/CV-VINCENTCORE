package userModel

type (
	RequestUser struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	RequestUserLogin struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)


