package model

import (
	"time"
)

type Person struct {
	Id        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
