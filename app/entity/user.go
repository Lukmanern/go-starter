package entity

import (
	"time"
)

type User struct {
	ID        uint       `gorm:"type:bigint(20) unsigned not null;autoIncrement;primaryKey" json:"id"`
	RoleID    uint       ``
	Role      Role       ``
	Name      string     `gorm:"type:varchar(20) not null"`
	Email     string     `gorm:"type:varchar(20) not null unique"`
	CreatedAt *time.Time `gorm:"type:timestamp null; default:null" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:timestamp null; default:null" json:"updated_at"`
	DeletedAt *time.Time `gorm:"type:timestamp null; default:null" json:"deleted_at"`
}

func (User) TableName() string {
	return "users"
}
