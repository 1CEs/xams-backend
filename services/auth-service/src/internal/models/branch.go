package models

type Branch struct {
	BranchID    uint    `json:"branch_id" gorm:"primaryKey;autoIncrement"`
	FacultyID   uint    `json:"faculty_id"`
	Faculty     Faculty `gorm:"constraint"`
	BranchName  string  `json:"branch_name" gorm:"unique;not null;varchar(255)"`
}