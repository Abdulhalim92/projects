package model

import "time"

type Student struct {
	StudentID int        `json:"student_id,omitempty" gorm:"primaryKey"`
	Name      string     `json:"name,omitempty"`
	Class     string     `json:"class,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type StudentDailyRecord struct {
	RecordID   int        `json:"record_id,omitempty" gorm:"primaryKey"`
	StudentID  int        `json:"student_id,omitempty"`
	SubjectID  int        `json:"subject_id,omitempty"`
	Attendance bool       `json:"attendance,omitempty"`
	Grade      float32    `json:"grade,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
}

type Subject struct {
	SubjectID int        `json:"subject_id,omitempty" gorm:"primaryKey"`
	Name      string     `json:"name,omitempty"`
	TeacherID int        `json:"teacher_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
