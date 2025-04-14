package entities

type CreateUserDto struct {
	FullName string	`json:"full_name"`
	Email string `json:"email"`
	Password string `json:"password"`
	City string `json:"city,omitempty"`
}
