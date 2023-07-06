package entity

import "time"

type Role struct {
	ID          uint       `gorm:"type:bigint(20) unsigned not null;autoIncrement;primaryKey"`
	Permissions Permission `gorm:"many2many:role_has_permissions"`
	Name        string     `gorm:"varchar(100) not null; unique;"`
	CreatedAt   *time.Time `gorm:"type:timestamp null; default:null"`
	UpdatedAt   *time.Time `gorm:"type:timestamp null; default:null"`
}
