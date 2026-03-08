package user

import "time"

type User struct {
	Id         int
	Username   string
	Occupation string
	Email      string
	Password   string
	AvatarUrl  string
	Role       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
