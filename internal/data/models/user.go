package models

import "github.com/osamikoyo/IM-auth/pkg/pb"

type User struct{
	ID uint64 `gorm:"primaryKey;autoIncrement"`
	Username string
	Email string
	Password string
}

func ToModels(u *pb.User) *User {
	return &User{
		ID: u.ID,
		Username: u.Username,
		Password: u.Password,
		Email: u.Email,
	}
}

func ToProtoBuf(u *User) *pb.User {
	return &pb.User{
		ID: u.ID,
		Email: u.Email,
		Password: u.Password,
		Username: u.Username,
	}
}