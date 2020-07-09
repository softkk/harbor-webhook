package database

import (
	"time"
)

//Project - Tables
type Project struct {
	// gorm.Model
	ID        int        `gorm:"primary_key"`
	Name      string     `gorm:"type:varchar(255);not null;unique"`
	Owner     int        `gorm:"type:int(11);not null"`
	isVolume  int        `gorm:"type:tinyint(1);not null;default:0"`
	Status    int        `gorm:"type:tinyint(1);not null;default:0"`
	CreatedAt *time.Time `gorm:"type:datetime;not null;column:createdAt"`
	UpdatedAt *time.Time `gorm:"type:datetime;not null;column:updatedAt"`
}

//SelectByName -
func (project *Project) SelectByName(name string) {
	MysqlDB.Where("name=?", name).First(project)
}
