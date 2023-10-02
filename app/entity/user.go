package entity

import "time"

type User struct {
	Timestampable

	email     string
	password  string
	lastLogin time.Time
}
