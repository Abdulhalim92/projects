package model

import "time"

type Record struct {
	RecordID  int
	Patientid int
	Doctorid  int
	Illnessid int
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
type Illness struct {
	ID   int
	name string
}
