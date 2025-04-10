package entities

type GetUserDto struct {
	ID int `json:"id"`
	FullName string	`json:"full_name"`
	Email string `json:"email"`
	City *string `json:"city,omitempty"`
}
