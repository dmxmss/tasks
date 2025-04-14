package entities

type User struct {
	ID int `gorm:"primaryKey"`
	FullName string `gorm:"not null"`
	Email string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	City string
}
