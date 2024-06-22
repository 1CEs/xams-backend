package models

type Faculty struct {
	FacultyID   uint   `json:"faculty_id" gorm:"primaryKey;autoIncrement"`
	FacultyName string `json:"faculty_name" gorm:"unique;not null;varchar(255)"`
}
