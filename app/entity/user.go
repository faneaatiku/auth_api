package entity

import "time"

const (
	StatusInactive UserStatus = 0
	StatusActive   UserStatus = 10
)

type UserStatus int

func (us UserStatus) IsActive() bool {
	return us == StatusActive
}

type User struct {
	Id             int64
	Email          string
	CanonicalEmail string
	Status         UserStatus
	Password       string
	Salt           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
