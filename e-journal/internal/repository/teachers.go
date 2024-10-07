package repository

import (
	"fmt"
	"gorm.io/gorm/clause"
	"log"
	"projects/internal/model"
)

func (r *Repository) GetTeachers() ([]model.Teacher, error) {
	var teachers []model.Teacher
	// select * from teachers
	err := r.db.Find(&teachers).Error
	if err != nil {
		log.Printf("GetTeachers: Error getting teachers: %v", err)
		return nil, err
	}
	if len(teachers) == 0 {
		log.Printf("GetTeachers: No teachers found")
		return nil, fmt.Errorf("GetTeachers: No teachers found")
	}

	return teachers, nil
}

func (r *Repository) GetTeacherByID(teacherID int) (*model.Teacher, error) {
	var teacher *model.Teacher
	// select * from teachers where teacher_id = ?
	err := r.db.First(&teacher, teacherID).Error
	if err != nil {
		log.Printf("GetTeacherByID: Error getting teacher by ID: %v", err)
		return nil, err
	}
	return teacher, nil
}
func (r *Repository) CreateTeacher(teacher *model.Teacher) (int, error) {
	// insert into teachers (name, user_id) values ('admin', 1) returning teacher_id
	err := r.db.Create(teacher).Error
	if err != nil {
		log.Printf("CreateTeacher: Error creating teacher: %v", err)

	}

	return teacher.TeacherID, nil
}
func (r *Repository) UpdateTeacher(teacher *model.Teacher) (*model.Teacher, error) {
	// update teachers set name = 'admin' where teacher_id = 1
	err := r.db.Clauses(clause.Returning{}).Updates(teacher).Error
	if err != nil {
		log.Printf("UpdateTeacher: Error updating teacher: %v", err)
		return nil, err
	}

	return teacher, nil

}
func (r *Repository) DeleteTeacher(teacherID int) error {
	// delete from teachers where teacher_id = 1
	err := r.db.Delete(&model.Teacher{}, teacherID).Error
	if err != nil {
		log.Printf("DeleteTeacher: Error deleting teacher: %v", err)
		return err
	}
	return nil
}
