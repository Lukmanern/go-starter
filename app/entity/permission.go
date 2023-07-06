package entity

import "time"

type Permission struct {
	ID        uint       `gorm:"type:bigint(20) unsigned not null;autoIncrement;primaryKey"`
	Name      string     `gorm:"varchar(100) not null; unique;"`
	CreatedAt *time.Time `gorm:"type:timestamp null; default:null"`
	UpdatedAt *time.Time `gorm:"type:timestamp null; default:null"`
}
