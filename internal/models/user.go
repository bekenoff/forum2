package models

type UserRegistration struct {
	Username string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserRegistration) Validate() bool {
	return false
}
