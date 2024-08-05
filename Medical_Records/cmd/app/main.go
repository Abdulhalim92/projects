package main

import (
	"fmt"
	"log"
	"projects/Medical_Records/internal/model"
	"projects/Medical_Records/internal/record"
	"time"
)

func main() {
	records := make([]model.Record, 0)
	NewRecords := record.NewRecords(records)
	NewService := record.NewService(*NewRecords)
	p1, err := NewService.CreatePatient(1, "Said", "Mirzoev", 20)
	if err != nil {
		log.Fatal(err)
	}
	d1, err := NewService.CreateDoctor(1, "Mister", "Doctor", "Surgery")
	if err != nil {
		log.Fatal(err)
	}
	_, err = NewService.CreateRecord(1, *p1, *d1, "ORVI", time.Date(1992, time.Month(12), 23, 11, 30, 30, 23, time.Local), time.Date(1993, time.Month(11), 2, 19, 34, 31, 6, time.Local))
	if err != nil {
		log.Fatal(err)
	}
	p2, err := NewService.CreatePatient(2, "John", "John", 4)
	if err != nil {
		log.Fatal(err)
	}
	d2, err := NewService.CreateDoctor(2, "Jack", "Jackson", "Therapist")
	if err != nil {
		log.Fatal(err)
	}
	_, err = NewService.CreateRecord(2, *p2, *d2, "Astma", time.Date(1998, time.Month(11), 11, 13, 36, 22, 1, time.Local), time.Date(1998, time.Month(14), 6, 43, 22, 1, 4, time.Local))
	if err != nil {
		log.Fatal(err)
	}
	p3, err := NewService.CreatePatient(3, "Rtr", "deafe", 34)
	if err != nil {
		log.Fatal(err)
	}
	_, err = NewService.CreateRecord(3, *p3, *d2, "cold", time.Date(2001, time.Month(10), 21, 8, 23, 3, 21, time.Local), time.Date(2002, time.Month(1), 5, 14, 3, 32, 7, time.Local))
	if err != nil {
		log.Fatal(err)
	}
	arr := NewService.ListAllRecords()
	fmt.Println(arr)
	r := NewService.ListRecordById(2)
	fmt.Println(*r)
	r2 := NewService.ListRecordsByDoctorId(2)
	fmt.Println(r2)
	r3 := NewService.ListRecordByPatientid(1)
	fmt.Println(*r3)
	result := NewService.ChangePatient(3, "Imran", "Khan", 34)
	fmt.Println(result)
	NewDoctor, err := NewService.CreateDoctor(4, "Jerax", "Alex", "Neurosurgery")
	if err != nil {
		log.Fatal(err)
	}
	NewDoctor2, err := NewService.CreateDoctor(5, "Mark", "Jack", "Pschychologist")
	if err != nil {
		log.Fatal(err)
	}
	result2 := NewService.ChangeDoctorInOneRecord(3, *NewDoctor2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result2)
	result3 := NewService.ChangeDoctorInAllRecords(2, *NewDoctor)
	fmt.Println(result3)
}
