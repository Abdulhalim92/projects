package model

import "time"

type Record struct {
	RecordID int
	Patient   Patient
	Doctor    Doctor
	Illness   string
	StartDate time.Time
	EndDate   time.Time
}
type Patient struct {
	ID      int
	Name    string
	Surname string
	Age     int
}
type Doctor struct {
	ID           int
	Name         string
	Surname      string
	TypeOfDoctor string
}
