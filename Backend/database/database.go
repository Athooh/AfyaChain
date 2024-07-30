package database

import (
	"log"
	"time"

	"github.com/Athooh/HealthChain/models"
	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDatabase() (db *gorm.DB, err error) {
	// Database connection
	dsn := "new_username:new_password@tcp(127.0.0.1:3306)/afya_chain_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
func CreatePatient(firstName, lastName string, dob time.Time, gender, email, phone, address string) *models.Patient {
	patient := &models.Patient{
		FirstName: firstName,
		LastName:  lastName,
		DOB:       dob,
		Gender:    gender,
		Email:     email,
		Phone:     phone,
		Address:   address,
	}
	db, err := ConnectDatabase()
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return nil
	}
	result := db.Create(patient)
	if result.Error != nil {
		log.Printf("Error creating patient: %v", result.Error)
		return nil
	}
	return patient
}

func GetPatient(id int) *models.Patient {
	var patient models.Patient
	db, err := ConnectDatabase()
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return nil
	}
	result := db.Preload("MedicalRecords").First(&patient, id)
	if result.Error != nil {
		log.Printf("Error retrieving patient: %v", result.Error)
		return nil
	}
	return &patient
}

func UpdatePatient(id int, firstName, lastName string, dob time.Time, gender, email, phone, address string) *models.Patient {
	patient := GetPatient(id)
	db, err := ConnectDatabase()
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return nil
	}
	if patient == nil {
		return nil
	}
	patient.FirstName = firstName
	patient.LastName = lastName
	patient.DOB = dob
	patient.Gender = gender
	patient.Email = email
	patient.Phone = phone
	patient.Address = address
	result := db.Save(patient)
	if result.Error != nil {
		log.Printf("Error updating patient: %v", result.Error)
		return nil
	}
	return patient
}

func DeletePatient(id int) bool {
	db, err := ConnectDatabase()
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return false
	}
	result := db.Delete(&models.Patient{}, id)
	if result.Error != nil {
		log.Printf("Error deleting patient: %v", result.Error)
		return false
	}
	return true
}

// CRUD functions for MedicalRecord

func CreateMedicalRecord(patientID int, recordDate time.Time, condition, treatment, notes string) *models.MedicalRecord {

	medicalRecord := &models.MedicalRecord{
		PatientID:  patientID,
		RecordDate: recordDate,
		Condition:  condition,
		Treatment:  treatment,
		Notes:      notes,
	}
	result := db.Create(medicalRecord)
	if result.Error != nil {
		log.Printf("Error creating medical record: %v", result.Error)
		return nil
	}
	return medicalRecord
}

func GetMedicalRecord(id int) *models.MedicalRecord {
	var medicalRecord models.MedicalRecord
	result := db.First(&medicalRecord, id)
	if result.Error != nil {
		log.Printf("Error retrieving medical record: %v", result.Error)
		return nil
	}
	return &medicalRecord
}

func UpdateMedicalRecord(id int, recordDate time.Time, condition, treatment, notes string) *models.MedicalRecord {
	medicalRecord := GetMedicalRecord(id)
	if medicalRecord == nil {
		return nil
	}
	medicalRecord.RecordDate = recordDate
	medicalRecord.Condition = condition
	medicalRecord.Treatment = treatment
	medicalRecord.Notes = notes
	result := db.Save(medicalRecord)
	if result.Error != nil {
		log.Printf("Error updating medical record: %v", result.Error)
		return nil
	}
	return medicalRecord
}

func DeleteMedicalRecord(id int) bool {
	result := db.Delete(&models.MedicalRecord{}, id)
	if result.Error != nil {
		log.Printf("Error deleting medical record: %v", result.Error)
		return false
	}
	return true
}
