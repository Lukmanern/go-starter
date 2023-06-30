package model

import "time"

type TodoCreate struct {
	Name      string `validate:"required"`
	Email     string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TodoUpdate struct {
	ID        uint
	Name      string `validate:"required"`
	Email     string `validate:"required"`
	UpdatedAt time.Time
}
