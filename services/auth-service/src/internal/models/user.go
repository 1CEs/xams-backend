package models

import "time"

type (

	Role string

	User struct {
		UserID    string    `json:"user_id" gorm:"primaryKey"`
		Password  string    `json:"password" gorm:"unique;not null;varchar(13)"`
		Email     string    `json:"email" gorm:"unique;not null;varchar(255)"`
		Prename   string    `json:"prename" gorm:"not null;varchar(50)"`
		FirstName string    `json:"first_name" gorm:"not null;varchar(255)"`
		LastName  string    `json:"last_name" gorm:"not null;varchar(255)"`
		BranchID  uint      `json:"branch_id" gorm:"not null;int"`
		Branch    Branch    `gorm:"constraint"`
		Role      Role   	`json:"role" gorm:"not null;type:enum('student', 'teacher', 'admin')"`
		CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
		UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	}
)

const (
	Student Role = "student"
	Teacher Role = "teacher"
	Admin Role = "admin"
)