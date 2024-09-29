package repository

import (
	"fmt"
	"log"
	"projects/internal/model"
)

func (r *Repository) CreateStudent(student *model.Student) (int, error) {
	// insert into students (name, class) values ('admin', 'admin') returning student_id
	err := r.db.Create(student).Error
	if err != nil {
		{
			log.Printf("CreateStudent: Error creating student: %v", err)
		}
		return 0, err
	}
	return student.StudentID, nil
}
func (r *Repository) GetStudentByID(studentID int) (*model.Student, error) {
	var student *model.Student
	// select * from students where student_id = ?
	err := r.db.First(&student, studentID).Error
	if err != nil {
		log.Printf("GetStudentByID: Error getting student by ID: %v", err)
		return nil, err
	}
	return student, nil
}
func (r *Repository) GetStudents() ([]model.Student, error) {
	var students []model.Student
	// select * from students
	err := r.db.Find(&students).Error
	if err != nil {
		log.Printf("GetStudents: Error getting students: %v", err)
		return nil, err
	}
	if len(students) == 0 {
		log.Printf("GetStudents: No students found")
		return nil, fmt.Errorf("GetStudents: No students found")
	}
	return students, nil
}
func (r *Repository) DeleteStudent(studentID int) error {
	// delete from students where student_id = 1
	err := r.db.Delete(&model.Student{}, studentID).Error
	if err != nil {
		log.Printf("DeleteStudent: Error deleting student: %v", err)
		return err
	}
	return nil
}
func (r *Repository) UpdateStudent(student *model.Student) error {
	// update students set name = 'admin', class = 'admin' where student_id = 1
	err := r.db.Save(student).Error
	if err != nil {
		log.Printf("UpdateStudent: Error updating student: %v", err)
		return err
	}
	return nil

}
