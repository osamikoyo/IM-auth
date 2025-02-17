package models

type User struct{
	ID uint `gorm:"primaryKey;autoIncrement"`
	Username string
	Email string
	Password string
}

