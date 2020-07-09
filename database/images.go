package database

import (
	"time"
)

//Image - Table Name
type Image struct {
	// gorm.Model
	ID        int        `gorm:"primary_key"`
	Name      string     `gorm:"type:varchar(255);not null"`
	Tag       string     `gorm:"type:varchar(255);not null"`
	RealPath  string     `gorm:"type:varchar(255);column:real_path"`
	Purpose   string     `gorm:"type:text"`
	Status    int        `gorm:"type:int(2);not null"`
	Applicant int        `gorm:"type:int(11)"`
	Review    int        `gorm:"type:int(11)"`
	CreatedAt *time.Time `gorm:"type:datetime;not null;column:createdAt"`
	UpdatedAt *time.Time `gorm:"type:datetime;not null;column:updatedAt"`
	// DeletedAt int        `gorm:"type:datetime;not null;column:deleted_at"`

	// Age      sql.NullInt64
	// IgnoreMe int `gorm:"-"` // ignore this field
}

//SelectByID -
func (image *Image) SelectByID(id int) {
	MysqlDB.Where("id=?", id).First(image)
}

//insert into k8s_test.`images` (`name`,`tag`, `real_path`, `purpose`, `status`, `applicant`, `review`        , `createdAt`, `updatedAt`)
//values ('iampserver','v1.0.5','chbcld.cminl.oa/a-line/iampserver:v1.0.5','CICD API更新',99,{專案管理人UID},1,NOW(),NOW())

//Insert -
func (image *Image) Insert() {

	// result := MysqlDB.Create(&image)

}
