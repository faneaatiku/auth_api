package entity

import "time"

type Timestampable struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}
