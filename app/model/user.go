package model

import "time"

type UserCreate struct {
	Name      string `validate:"required"`
	Email     string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserUpdate struct {
	ID        uint
	Name      string `validate:"required"`
	Email     string `validate:"required"`
	UpdatedAt time.Time
}
