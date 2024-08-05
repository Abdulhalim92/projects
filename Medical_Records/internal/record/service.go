package record

import (
	"projects/Medical_Records/internal/model"
	"time"
)

type Service struct {
	Records Records
}

func NewService(r Records) *Service {
	return &Service{r}
}
func (s *Service) CreatePatient(ID int, Name, Surname string, Age int) (*model.Patient, error) {
	return s.Records.NewPatient(ID, Name, Surname, Age)
}
func (s *Service) CreateDoctor(ID int, Name, Surname string, TypeOfDoctor string) (*model.Doctor, error) {
	return s.Records.NewDoctor(ID, Name, Surname, TypeOfDoctor)
}
func (s *Service) CreateRecord(id int, p model.Patient, d model.Doctor, illness string, starttime time.Time, endtime time.Time) (*model.Record, error) {
	return s.Records.AddRecord(id, p, d, illness, starttime, endtime)
}
func (s *Service) ListAllRecords() []model.Record {
	return s.Records.GetAllRecords()
}
func (s *Service) ListRecordById(id int) *model.Record {
	return s.Records.GetRecordByRecordID(id)
}
func (s *Service) ListRecordsByDoctorId(id int) []model.Record {
	return s.Records.GetRecordsByDoctorID(id)
}
func (s *Service) ListRecordByPatientid(id int) *model.Record {
	return s.Records.GetRecordByPatientID(id)
}
func (s *Service) ChangePatient(id int, name, surname string, age int) bool {
	return s.Records.UpdatePatient(id, name, surname, age)
}
func (s *Service) ChangeDoctorInAllRecords(id int, NewDoctor model.Doctor) bool {
	return s.Records.UpdateDoctorInAllRecords(id, NewDoctor)
}
func (s *Service) ChangeDoctorInOneRecord(patientid int, NewDoctor model.Doctor) bool {
	return s.Records.UpdateDoctorInOneRecord(patientid, NewDoctor)
}
func (s *Service) RemoveRecordByPatientID(patientid int) bool {
	return s.Records.DeleteRecordByPatientID(patientid)
}
func (s *Service) RemoveRecordByRecordID(recordid int) bool {
	return s.Records.DeleteRecordByRecordID(recordid)
}
