package database

import "time"

// ProjectImage - Table
type ProjectImage struct {
	ImageID   int        `gorm:"primary_key;column:imageId"`
	ProjectID int        `gorm:"primary_key;column:projectId"`
	CreatedAt *time.Time `gorm:"type:datetime;not null;column:createdAt"`
	UpdatedAt *time.Time `gorm:"type:datetime;not null;column:updatedAt"`
}

//TableName - 自訂Table name, gorm在對應的結構體默認會加個s
func (projectImage ProjectImage) TableName() string {
	return "project_image"
}
