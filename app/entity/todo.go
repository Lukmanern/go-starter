package entity

import (
	"time"
)

type status = string

const (
	DONE    status = "done"
	TODO    status = "todo"
	DELETED status = "deleted"
)

type Todo struct {
	ID        uint       `gorm:"type:bigint(20) unsigned not null;autoIncrement;primaryKey"`
	UserID    uint       ``
	User      User       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Todo      string     `gorm:"type:text not null"`
	Status    status     `gorm:"type:enum('done', 'todo', 'deleted')"`
	CreatedAt *time.Time `gorm:"type:timestamp null; default:null"`
	UpdatedAt *time.Time `gorm:"type:timestamp null; default:null"`
	DeletedAt *time.Time `gorm:"type:timestamp null; default:null"`
}

func (Todo) TableName() string {
	return "todos"
}
