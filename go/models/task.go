package models

type Task struct {
	Id       uint
	UserID   int
	MasterID int
	Person   string
	Date     string
}
