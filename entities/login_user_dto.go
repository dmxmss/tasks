package entities

type LoginUserDto struct {
	FullName string `json:"full_name"`
	Email string `json:"email"`
	Password string `json:"password"`
}
