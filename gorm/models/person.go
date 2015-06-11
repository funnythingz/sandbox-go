package models

import (
	"time"
)

type Person struct {
	Id        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
