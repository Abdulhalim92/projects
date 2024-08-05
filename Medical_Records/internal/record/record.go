package record

import (
	"errors"
	"fmt"
	"projects/Medical_Records/internal/model"
	"time"
)

type Records struct {
	Records []model.Record
}

var LastRecordID, LastPatientID int

func init() {
	LastRecordID = 0
}
func (r *Records) NewPatient(ID int, Name, Surname string, Age int) (*model.Patient, error) {
	for _, v := range r.Records {
		if v.Patient.ID == ID {
			return nil, errors.New("such id exists")
		}
	}
	return &model.Patient{ID: ID, Name: Name, Surname: Surname, Age: Age}, nil
}
func (r *Records) NewDoctor(ID int, Name, Surname string, TypeOfDoctor string) (*model.Doctor, error) {
	for _, v := range r.Records {
		if v.Doctor.ID == ID {
			return nil, errors.New("such doctor exists")
		}
	}
	return &model.Doctor{ID: ID, Name: Name, Surname: Surname, TypeOfDoctor: TypeOfDoctor}, nil
}
func NewRecords(records []model.Record) *Records {
	return &Records{
		Records: records,
	}
}
func (r *Records) AddRecord(id int, p model.Patient, d model.Doctor, illness string, starttime time.Time, endtime time.Time) (*model.Record, error) {
	for _, v := range r.Records {
		if v.RecordID == id {
			return nil, errors.New("such RecordID exists")
		} else if v.Patient.ID == id {
			return nil, errors.New("such PatientID exists")
		}

	}
	LastRecordID++

	record := model.Record{RecordID: id, Patient: p, Doctor: d, Illness: illness, StartDate: starttime, EndDate: endtime}
	r.Records = append(r.Records, record)
	fmt.Printf("Record with doctor '%v' and Patient '%v' created, id '%d' , illness '%s',  starttime '%v',  endtime '%v' \n", d, p, id, illness, starttime, endtime)
	return &record, nil
}
func (r Records) GetAllRecords() []model.Record {
	return r.Records
}
func (r *Records) GetRecordByRecordID(id int) *model.Record {
	for _, r := range r.Records {
		if r.RecordID == id {
			return &r
		}
	}
	return nil
}
func (r Records) GetRecordsByDoctorID(id int) []model.Record {
	x := make([]model.Record, 0)
	for _, v := range r.Records {
		if v.Doctor.ID == id {
			x = append(x, v)
		}
	}
	return x
}
func (r *Records) GetRecordByPatientID(id int) *model.Record {
	for _, v := range r.Records {
		if v.Patient.ID == id {
			return &v
		}
	}
	return nil
}
func (r *Records) UpdatePatient(id int, name, surname string, age int) bool {
	for i := 0; i < len(r.Records); i++ {
		if r.Records[i].Patient.ID == id {
			r.Records[i].Patient = model.Patient{
				ID:      id,
				Name:    name,
				Surname: surname,
				Age:     age,
			}
			return true
		}
	}
	return false
}
func (r *Records) UpdateDoctorInAllRecords(id int, NewDoctor model.Doctor) bool {
	var j int
	for i := 0; i < len(r.Records); i++ {
		if r.Records[i].Doctor.ID == id {
			r.Records[i].Doctor = model.Doctor{
				ID:           id,
				Name:         NewDoctor.Name,
				Surname:      NewDoctor.Surname,
				TypeOfDoctor: NewDoctor.TypeOfDoctor,
			}
			j++
		}
	}
	return j != 0
}
func (r *Records) UpdateDoctorInOneRecord(PatientID int, NewDoctor model.Doctor) bool {
	for i := 0; i < len(r.Records); i++ {
		if r.Records[i].Patient.ID == PatientID {
			r.Records[i].Doctor = model.Doctor{
				ID:           NewDoctor.ID,
				Name:         NewDoctor.Name,
				Surname:      NewDoctor.Surname,
				TypeOfDoctor: NewDoctor.TypeOfDoctor,
			}
			return true
		}
	}
	return false
}
func (r *Records) DeleteRecordByPatientID(PatientID int) bool {
	for i := 0; i < len(r.Records); i++ {
		if r.Records[i].Patient.ID == PatientID {
			r.Records = append(r.Records[:i], r.Records[i+1:]...)
			return true
		}
	}
	return false
}
func (r *Records) DeleteRecordByRecordID(RecordID int) bool {
	for i := 0; i < len(r.Records); i++ {
		if r.Records[i].RecordID == RecordID {
			r.Records = append(r.Records[:i], r.Records[i+1:]...)
			return true
		}
	}
	return false
}
