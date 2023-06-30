package entity

import (
	"time"
)

type Todo struct {
	ID        uint       `gorm:"type:bigint(20) unsigned not null;autoIncrement;primaryKey" `
	UserID    uint       ``
	User      User       ``
	Todo      string     `gorm:"type:text not null" `
	CreatedAt *time.Time `gorm:"type:timestamp null; default:null" `
	UpdatedAt *time.Time `gorm:"type:timestamp null; default:null" `
	DeletedAt *time.Time `gorm:"type:timestamp null; default:null" `
}

func (Todo) TableName() string {
	return "todos"
}

type User struct {
	ID        uint `gorm:"type:bigint(20) unsigned not null;autoIncrement;primaryKey" `
	Todo      []Todo
	Name      string     `gorm:"type:varchar(20) not null"`
	Email     string     `gorm:"type:varchar(20) not null unique"`
	CreatedAt *time.Time `gorm:"type:timestamp null; default:null" `
	UpdatedAt *time.Time `gorm:"type:timestamp null; default:null" `
	DeletedAt *time.Time `gorm:"type:timestamp null; default:null" `
}

func (User) TableName() string {
	return "users"
}
